#!/usr/bin/env python3

# Copyright 2026 Flant JSC
#
# Licensed under the Apache License, Version 2.0

"""
Detect final werf images that should be scanned for the current PR.

Algorithm:
  1. Get PR commits.
  2. Read current images_tags_werf.json.
  3. Keep only entries with:
       Final == true
       Commit in PR commits
  4. Convert werf image name into module/image fields for scanner output.
  5. Write changed_images.json and GitHub Actions outputs.
"""

import json
import os
import subprocess
import sys


def run(cmd: list[str]) -> str:
    return subprocess.check_output(cmd, text=True).strip()


def load_build_report(path: str) -> dict:
    with open(path) as fp:
        report = json.load(fp)

    images = report.get("Images")
    if isinstance(images, list):
        normalized = {}
        for entry in images:
            key = entry.get("WerfImageName") or entry.get("Name") or entry.get("Image")
            if key:
                normalized[key] = entry
        images = normalized

    if not isinstance(images, dict):
        raise SystemExit(f"unexpected build report shape at {path}: no Images map")

    return images


def get_pr_commits() -> set[str]:
    base_ref = os.environ.get("GITHUB_BASE_REF")
    if not base_ref:
        raise SystemExit("GITHUB_BASE_REF is not set")

    if not base_ref.startswith("origin/"):
        base_ref = f"origin/{base_ref}"

    merge_base = run(["git", "merge-base", base_ref, "HEAD"])
    commits = run(["git", "log", "--format=%H", f"{merge_base}..HEAD"])

    return set(commits.splitlines()) if commits else set()


def split_werf_image_name(name: str) -> tuple[str, str]:
    if "/" not in name:
        return name, name

    module, image = name.split("/", 1)
    return module, image


def compute_changed(images: dict, pr_commits: set[str]) -> list:
    changed = []

    for name, entry in images.items():
        if not isinstance(entry, dict):
            continue

        if entry.get("Final") is not True:
            continue

        commit = entry.get("Commit")
        if not commit or commit not in pr_commits:
            continue

        digest = entry.get("DockerImageDigest")
        if not digest:
            continue

        werf_image_name = entry.get("WerfImageName") or name
        module, image = split_werf_image_name(werf_image_name)

        changed.append({
            "module": module,
            "image": image,
            "digest": digest,
            "commit": commit,
            "werf_image_name": werf_image_name,
        })

    changed.sort(key=lambda c: (c["module"], c["image"]))
    return changed


def emit_github_outputs(changed: list) -> None:
    out_path = os.environ.get("GITHUB_OUTPUT")
    if not out_path:
        return

    matrix = {"include": changed}
    compact = [f"{c['module']}.{c['image']}" for c in changed]

    with open(out_path, "a") as fp:
        fp.write(f"changed_count={len(changed)}\n")
        fp.write(f"matrix={json.dumps(matrix, separators=(',', ':'))}\n")
        fp.write(f"changed_compact={json.dumps(compact, separators=(',', ':'))}\n")


def main() -> int:
    build_report_path = os.environ.get("BUILD_REPORT_PATH", "images_tags_werf.json")
    out_changed = os.environ.get("OUTPUT_CHANGED", "changed_images.json")

    if not os.path.exists(build_report_path):
        raise SystemExit(f"ERROR: build report not found: {build_report_path}")

    pr_commits = get_pr_commits()
    print(f"PR commits: {len(pr_commits)}")

    images = load_build_report(build_report_path)
    print(f"Build report total entries: {len(images)}")

    changed = compute_changed(images, pr_commits)

    with open(out_changed, "w") as fp:
        json.dump(changed, fp, indent=2)

    print(f"Changed final images: {len(changed)}")

    if changed:
        print("Images for scan:")
        for c in changed:
            print(f"  {c['werf_image_name']}  {c['commit']}  {c['digest']}")

    emit_github_outputs(changed)
    return 0


if __name__ == "__main__":
    sys.exit(main())
