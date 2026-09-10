# Conventions

## Conventional Commits

Use Conventional Commits for every commit.
Use `<type>(<scope>): <description>`.
Use an imperative summary after the colon.
Use `feat` for features and `fix` for fixes.
Use `docs`, `chore`, `refactor`, `style`, or `test` when appropriate.

## Git Workflow

ALWAYS work on a feature branch.
NEVER work directly on `main` or `master`.
Use `specs/state.yaml` as the workflow-mode signal.
Use `solo-git` workflow mode unless the user selects team PR workflow.

## Always Green / Shift Left

Keep Preflight and CI green before forward work.
Run `task preflight` as Preflight.
Run `gh pr checks` when a pull request exists.
Fix failures early because integration costs more than local fixes.

## Discovered Defects

Treat every reproducible gate failure as a discovered defect.
Use `quick-fix` for a trivial data-only fix.
Use `fix-bug` when investigation or TDD is necessary.
Write a bug specification when reproduction remains blocked.
DO NOT continue forward work on a red gate.
Keep discovered fixes in separate Conventional Commits.

## Banned Dismissals

DO NOT call a reproducible failure pre-existing.
DO NOT call a reproducible failure unrelated.
DO NOT call a reproducible failure out of scope.
Fix or log the failure instead.

## Go Code

Use `gofmt` through the existing toolchain.
Use Go doc comments for exported APIs.
Keep tests beside the code they test.
Use `testData/` for golden fixtures.
Use functional options for optional library behavior.
Return and wrap errors with useful context.
Keep filesystem I/O in the CLI layer.

## Lint and Dependencies

NEVER modify `.golangci.yml`.
ALWAYS fix lint errors in the codebase.
Ignore lint errors only when absolutely necessary.
NEVER modify `vendor/`.
Run `go mod vendor` when vendor regeneration is necessary.
Notify the user when a vendor fix is necessary.
Use a workaround or report the blocker when an upstream fix cannot happen.

## Bigpowers Scripts

Use full bigpowers script paths under `~/.pi/agent/npm/node_modules/bigpowers/`.
DO NOT rely on the current working directory for bigpowers script paths.

## Defensive Code

Use normal Go error handling for this local documentation generator.
Add defensive network behavior only when introducing external network calls.

## Planning Output

Write planning and verification output under `specs/`.
Update `specs/tech-architecture/tech-stack.md` after architecture changes.
Use `specs/bugs/` for bug investigations.
Use `specs/verifications/` for verification evidence.
