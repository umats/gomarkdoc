# Security Review: GitLab output format

## Scope

`format/gitlab.go` and CLI format selection.

## Findings

No HIGH-confidence security findings.

The change formats repository URLs already detected or explicitly supplied by the user. It makes no network requests and adds no authentication or secret handling.
