# e01s01: Add GitLab output format

## 1. Story ID

e01s01

## 2. Title

Add GitLab output format.

## 3. Maturity

3 — Countable.

## 4. Type

Feature.

## 5. Business Narrative

As a Go documentation publisher using GitLab, I want GitLab-compatible output so links render correctly without template overrides.

## 6. Problem

The existing GitHub output uses different heading-link and source-link rules.

## 7. Actors

CLI users publishing generated Markdown on GitLab.com or self-hosted GitLab.

## 8. Preconditions

A repository URL is detected or supplied through the existing repository options.

## 9. Main Flow

The user selects `--format gitlab`; gomarkdoc renders GitLab-compatible anchors and source links.

## 10. Alternative Flows

Without repository metadata, source links remain empty, as with existing formats.

## 11. Requirements

### ADDED: GitLab output format

gomarkdoc provides a `gitlab` format selection that renders GitLab-compatible heading links and source links.

## 12. Constraints

Reuse `format.Format` and existing repository metadata. Add no dependency.

## 13. Dependencies

None.

## 14. Interfaces

The existing `format.Format` interface and CLI `--format` value.

## 15. Data

Existing `lang.Location` and `lang.Repo` values.

## 16. Error Handling

An unknown format continues to return the existing invalid-format error.

## 17. Acceptance Criteria

```gherkin
Scenario: Generate GitLab Markdown
  Given a repository URL for GitLab.com or a self-hosted GitLab instance
  When the user selects the gitlab format
  Then source links use /-/blob/
  And generated anchors follow GitLab rules

Scenario: Preserve existing formats
  Given the supported GitHub, Azure DevOps, and plain formats
  When GitLab support is added
  Then their existing tests pass unchanged
```

Covers: SC-e01s01-P1-01, SC-e01s01-P1-02, SC-e01s01-P1-03.

## 18. Out of Scope

Remote-based format detection, GitLab APIs, and changes to existing formats.

## 19. Risks

The command fixture harness must include GitLab without weakening existing golden checks.

## 20. Verification Script

1. Run `go test ./format -run '^TestGitLabFlavoredMarkdown'`.
2. Run `go test ./cmd/gomarkdoc`.
3. Run `task preflight`.
