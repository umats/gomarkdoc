# Impact: GitLab output format

## Target

`format` package and `cmd/gomarkdoc/command.go` format selection.

## Dependents

- `cmd/gomarkdoc/command.go`: constructs the selected `format.Format` implementation.
- `renderer.go` and templates: consume `format.Format` through its existing contract.
- `format/github_test.go` and `cmd/gomarkdoc/command_test.go`: establish the closest behavior and integration coverage.

## Affected Stories

- e01s01: Add GitLab output format.

## Test Coverage

- `format/github_test.go`: formatter behavior and repository source-link expectations.
- `cmd/gomarkdoc/command_test.go`: runs every supported output format through generated-file fixtures.
- Gap: no GitLab formatter or CLI fixture exists.

## Risk: Medium

The extension uses a stable interface, but it is user-visible and the command harness must include the new format without changing existing output.

## Recommended action

Add focused formatter and CLI integration tests, then add the format as a new implementation. Do not modify shared formatting behavior.
