# Test Design: e01 GitLab Markdown

## 1. Risk Matrix & Scenarios

| Scenario ID | Behavior | Risk | Level | Target |
| --- | --- | --- | --- | --- |
| SC-e01s01-P1-01 | GitLab anchors lowercase, remove non-word text, and collapse repeated hyphens. | P1 | Unit | `format/gitlab_test.go` |
| SC-e01s01-P1-02 | GitLab source links use the configured GitLab.com or self-hosted remote and `/-/blob/`. | P1 | Unit | `format/gitlab_test.go` |
| SC-e01s01-P1-03 | `--format gitlab` produces the GitLab fixture without changing other format fixtures. | P1 | Integration | `cmd/gomarkdoc/command_test.go` |

## 2. Fixture Architecture & Isolation

Use existing `lang.Location` values for source-link unit tests and the existing command test harness plus `testData` golden output. No network, API, or new fixture framework is needed.

## 3. NFR Verification

| Requirement | Verification |
| --- | --- |
| Existing outputs remain compatible. | `task preflight` |

## 4. Out of Scope

- Browser-rendered GitLab UI testing.
- GitLab API or authentication tests.
