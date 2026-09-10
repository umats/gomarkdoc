# gomarkdoc — Claude Code

Read `CONVENTIONS.md` before Git or GitHub operations.

<!-- BEGIN bigpowers:context-routing -->
## Context Routing

Read `specs/tech-architecture/tech-stack.md` for architecture decisions.
Read `specs/` before planning or implementation work.
<!-- END bigpowers:context-routing -->

<!-- BEGIN bigpowers:project -->
## Project

`gomarkdoc` formats Go package documentation as Markdown.
Stack: Go 1.26, Cobra, Viper, `text/template`, Task, and GoReleaser.

## Commands

| Action | Command |
| --- | --- |
| Run | `go run ./cmd/gomarkdoc --help` |
| Test | `task test` |
| Build | `task build` |
| Lint | `task lint` |
| Preflight | `task preflight` |
| CI | `gh pr checks` |

## Architecture

The CLI configures package discovery and parsing through `lang`.
The renderer applies templates and Markdown formatters before writing output.

## Conventions

- Use Conventional Commits.
- Use feature branches for every task.
- Add Go doc comments to exported APIs.
- Keep tests beside code.
- Use `testData/` for golden fixtures.
- Use functional options for optional library behavior.
- Use full bigpowers script paths under `~/.pi/agent/npm/node_modules/bigpowers/`.

## Never

- NEVER modify `.golangci.yml`.
- ALWAYS fix lint failures in code.
- Ignore lint failures only when absolutely necessary.
- NEVER modify `vendor/`.
- Run `go mod vendor` when vendor regeneration is necessary.
- Report vendor fixes to the user.
- Report an upstream dependency blocker when no safe workaround exists.

## Agent Rules

- Use bigpowers skills for structured work.
- Keep Preflight and CI green before forward work.
- Fix or log every reproducible gate failure.
- Put planning output in `specs/`.
- Write the minimum code that solves the stated problem.
- Run relevant tests after every change.
<!-- END bigpowers:project -->
