#!/usr/bin/env python3

# Copyright 2026 Flant JSC
#
# Licensed under the Apache License, Version 2.0

"""
Detect module images that should be scanned for the current PR.

Inputs (env):
  BUILD_REPORT_PATH
      Path to the current werf build report (images_tags_werf.json).

      The report contains all images produced during the build, including
      intermediate stages and final images. Each image records:
        - Final
        - Commit
        - DockerImageDigest
        - other build metadata.

  IMAGES_DIGESTS_PATH
      Path to images_digests.json extracted from the assembled dev image.

      This file maps module images to their digests:

          module.image -> digest

      It is used to convert digests from the build report into scanner
      image keys.

  OUTPUT_CHANGED
      Output file with module images selected for scanning.

  GITHUB_OUTPUT
      GitHub Actions outputs:
        - changed_count
        - matrix
        - changed_compact

Algorithm:

  1. Determine commits that belong to the current PR:

         git log --format=%H $(git merge-base origin/<base> HEAD)..HEAD

  2. Read the current build report.

  3. Keep only final images:

         Final == true

  4. For every final image, check its Commit.

     If Commit belongs to the current PR commit set, keep its
     DockerImageDigest.

     Commit is used instead of Rebuilt because Rebuilt only indicates
     whether werf rebuilt an image during the current run. Cached images
     produced by previous runs of the same PR may have Rebuilt=false while
     still belonging to the PR.

  5. Read images_digests.json.

     Match collected digests against module image digests to obtain scanner
     keys in the form:

         module.image

  6. Write changed_images.json and GitHub Actions outputs.

Empty result:

  If no module images are matched, changed_images.json is written as an
  empty array.
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


def collect_relevant_digests(images: dict, pr_commits: set[str]) -> dict:
    out = {}

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

        out[digest] = {
            "werf_image_name": name,
            "commit": commit,
        }

    return out


def compute_changed(images_digests: dict, relevant_digests: dict) -> list:
    changed = []

    for module, mod_images in (images_digests or {}).items():
        if not isinstance(mod_images, dict):
            continue

        for image, digest in mod_images.items():
            if not isinstance(digest, str):
                continue

            meta = relevant_digests.get(digest)
            if not meta:
                continue

            changed.append({
                "module": module,
                "image": image,
                "digest": digest,
                "commit": meta["commit"],
                "werf_image_name": meta["werf_image_name"],
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
    images_digests_path = os.environ.get("IMAGES_DIGESTS_PATH", "images_digests.json")
    out_changed = os.environ.get("OUTPUT_CHANGED", "changed_images.json")

    if not os.path.exists(build_report_path):
        raise SystemExit(f"ERROR: build report not found: {build_report_path}")

    if not os.path.exists(images_digests_path):
        raise SystemExit(f"ERROR: images digests not found: {images_digests_path}")

    pr_commits = get_pr_commits()
    print(f"PR commits: {len(pr_commits)}")

    images = load_build_report(build_report_path)
    print(f"Build report total entries: {len(images)}")

    relevant_digests = collect_relevant_digests(images, pr_commits)
    print(f"Final images with Commit from PR: {len(relevant_digests)}")

    with open(images_digests_path) as fp:
        images_digests = json.load(fp)

    changed = compute_changed(images_digests, relevant_digests)

    with open(out_changed, "w") as fp:
        json.dump(changed, fp, indent=2)

    print(f"Changed module images: {len(changed)}")

    if changed:
        print("Images for scan:")
        for c in changed:
            print(f"  {c['module']}.{c['image']}  {c['commit']}  {c['digest']}")

    emit_github_outputs(changed)
    return 0


if __name__ == "__main__":
    sys.exit(main())
