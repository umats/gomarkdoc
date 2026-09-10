---
bug_id: BUG-2026-09-10T103348Z
status: fixed
severity: medium
scope: package-discovery
title: Recursive discovery rejects test-only directories
---

# Recursive discovery rejects test-only directories

## Problem

Running `gomarkdoc ./...` fails when a recursive walk reaches a directory that contains only external test files. In `/home/umats/projects/issapi`, `tests/client` has only `*_test.go` files and fails with `no source-code package in directory tests/client`.

Expected behavior: recursive discovery skips test-only directories, as it already skips directories that cannot be loaded as packages.

Security impact: NONE — no security exploit path identified.

## Root Cause Analysis

Recursive discovery loads the directory successfully, then passes a build package with no non-test Go files to documentation parsing. The parser filters to source files and returns no package. The wildcard path does not treat this expected absence as skippable.

Risk: Medium — recursive generation is blocked for repositories that keep test-only packages.

## TDD Fix Plan

1. **RED:** Add a command-package test for a wildcard package specification containing only an external test file.
   **GREEN:** Skip wildcard build packages with no source or cgo files.
   **verify:** `go test ./cmd/gomarkdoc -run '^TestLoadPackagesSkipsTestOnlyPackage$'`

## Acceptance Criteria

- [x] Recursive discovery skips test-only directories.
- [x] Explicit source packages still load normally.
- [x] Existing tests pass.

## Resolution

Wildcard discovery now skips build packages with no source or cgo files. The regression test and `/home/umats/projects/issapi` reproduction pass.
