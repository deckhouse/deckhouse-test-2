#!/usr/bin/env python3

# Copyright 2026 Flant JSC
#
# Licensed under the Apache License, Version 2.0

"""
Detect final werf images that should be scanned for the current PR.

Inputs:
  BUILD_REPORT_PATH
      Path to images_tags_werf.json.

  OUTPUT_CHANGED
      Output file with selected images.
      Default: changed_images.json.

  GITHUB_BASE_REF
      Pull request base branch name.
      Example: main.

  GITHUB_OUTPUT
      GitHub Actions outputs:
        changed_count
        matrix
        changed_compact

Algorithm:
  1. Get commits that belong to the PR:
       git log --format=%H $(git merge-base origin/<base> HEAD)..HEAD

  2. Read images_tags_werf.json.

  3. Keep only final images:
       Final == true

  4. For each final image, check commit:
       Images[name].Commit
       Images[name].Stages[].Commit

  5. If any commit belongs to the PR, write scanner item:
       module
       image
       digest
       commit
       werf_image_name

  6. If build report is missing, write empty result and exit 0.
     This covers cases where build did not produce images.
"""

import json
import os
import subprocess
import sys
from typing import Any


def run(cmd: list[str], check: bool = True) -> str:
    try:
        return subprocess.check_output(
            cmd,
            text=True,
            stderr=subprocess.STDOUT,
        ).strip()
    except subprocess.CalledProcessError as e:
        if check:
            print(f"ERROR: command failed: {' '.join(cmd)}", file=sys.stderr)
            print(e.output, file=sys.stderr)
            raise
        return ""


def write_changed(path: str, changed: list[dict[str, Any]]) -> None:
    with open(path, "w") as fp:
        json.dump(changed, fp, indent=2)
        fp.write("\n")


def emit_github_outputs(changed: list[dict[str, Any]]) -> None:
    out_path = os.environ.get("GITHUB_OUTPUT")
    if not out_path:
        return

    matrix = {"include": changed}
    compact = [f"{c['module']}.{c['image']}" for c in changed]

    with open(out_path, "a") as fp:
        fp.write(f"changed_count={len(changed)}\n")
        fp.write(f"matrix={json.dumps(matrix, separators=(',', ':'))}\n")
        fp.write(f"changed_compact={json.dumps(compact, separators=(',', ':'))}\n")


def load_build_report(path: str) -> dict[str, Any]:
    with open(path) as fp:
        report = json.load(fp)

    images = report.get("Images")

    if isinstance(images, list):
        normalized = {}

        for entry in images:
            if not isinstance(entry, dict):
                continue

            key = (
                entry.get("WerfImageName")
                or entry.get("Name")
                or entry.get("Image")
            )

            if key:
                normalized[key] = entry

        images = normalized

    if not isinstance(images, dict):
        raise SystemExit(f"ERROR: unexpected build report shape at {path}: no Images map")

    return images


def get_base_ref() -> str:
    base_ref = os.environ.get("GITHUB_BASE_REF")

    if not base_ref:
        raise SystemExit("ERROR: GITHUB_BASE_REF is not set")

    if not base_ref.startswith("origin/"):
        base_ref = f"origin/{base_ref}"

    return base_ref


def print_git_debug(base_ref: str) -> None:
    print("Git debug:")
    print(f"  HEAD: {run(['git', 'rev-parse', 'HEAD'], check=False)}")
    print(f"  base_ref: {base_ref}")
    print("  branches:")
    branches = run(["git", "branch", "-avv"], check=False)
    for line in branches.splitlines():
        print(f"    {line}")

    print("  last commits:")
    commits = run(["git", "log", "--oneline", "-10"], check=False)
    for line in commits.splitlines():
        print(f"    {line}")


def get_pr_commits() -> set[str]:
    base_ref = get_base_ref()
    print_git_debug(base_ref)

    merge_base = run(["git", "merge-base", base_ref, "HEAD"])
    print(f"Merge base: {merge_base}")

    commits = run(["git", "log", "--format=%H", f"{merge_base}..HEAD"])

    pr_commits = set(commits.splitlines()) if commits else set()

    print(f"PR commits: {len(pr_commits)}")

    if pr_commits:
        print("PR commit list:")
        for commit in sorted(pr_commits):
            print(f"  {commit}")

    return pr_commits


def split_werf_image_name(name: str) -> tuple[str, str]:
    if "/" not in name:
        return name, name

    module, image = name.split("/", 1)
    return module, image


def find_pr_commit(entry: dict[str, Any], pr_commits: set[str]) -> str | None:
    commit = entry.get("Commit")

    if isinstance(commit, str) and commit in pr_commits:
        return commit

    stages = entry.get("Stages", [])

    if not isinstance(stages, list):
        return None

    for stage in stages:
        if not isinstance(stage, dict):
            continue

        stage_commit = stage.get("Commit")

        if isinstance(stage_commit, str) and stage_commit in pr_commits:
            return stage_commit

    return None


def compute_changed(
    images: dict[str, Any],
    pr_commits: set[str],
) -> list[dict[str, Any]]:
    changed = []

    for name, entry in images.items():
        if not isinstance(entry, dict):
            continue

        if entry.get("Final") is not True:
            continue

        digest = entry.get("DockerImageDigest")
        if not isinstance(digest, str) or not digest:
            continue

        werf_image_name = entry.get("WerfImageName") or name
        if not isinstance(werf_image_name, str) or not werf_image_name:
            continue

        commit = find_pr_commit(entry, pr_commits)
        if not commit:
            continue

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


def main() -> int:
    build_report_path = os.environ.get("BUILD_REPORT_PATH", "images_tags_werf.json")
    out_changed = os.environ.get("OUTPUT_CHANGED", "changed_images.json")

    if not os.path.exists(build_report_path):
        print(f"Build report not found: {build_report_path}")
        print("Write empty changed images and exit 0.")

        changed: list[dict[str, Any]] = []
        write_changed(out_changed, changed)
        emit_github_outputs(changed)

        return 0

    pr_commits = get_pr_commits()

    images = load_build_report(build_report_path)
    print(f"Build report total entries: {len(images)}")

    changed = compute_changed(images, pr_commits)

    write_changed(out_changed, changed)

    print(f"Changed final images: {len(changed)}")

    if changed:
        print("Images for scan:")
        for c in changed:
            print(
                f"  {c['werf_image_name']}  "
                f"{c['commit']}  "
                f"{c['digest']}"
            )
    else:
        print("No final images matched PR commits.")

    emit_github_outputs(changed)

    return 0


if __name__ == "__main__":
    sys.exit(main())
