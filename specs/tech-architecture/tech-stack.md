# Project Context

## Stack

- Go module: `github.com/umats/gomarkdoc`; `go.mod` declares Go 1.26.
- A library plus a Cobra CLI (`cmd/gomarkdoc`) that turns Go package documentation into Markdown.
- Standard-library analysis pipeline: `go/build` selects files, `go/parser` and `go/doc` produce documentation structures, and `go/ast` filters unexported declarations.
- Configuration: Viper reads `.gomarkdoc.*`, environment variables, and bound Cobra flags. CLI flags take precedence.
- Rendering: `text/template` templates over the `lang` model, with a small `format.Format` interface and GitHub, Azure DevOps, and plain Markdown implementations.
- Repository links: `go-git/v5` discovers repository metadata; CLI flags can override it.
- Console logging: Logrus with `logrus-prefixed-formatter`, sent to stderr and controlled by `-v` / `-vv`.
- Build and release: Task, GitHub Actions, GoReleaser. Releases are static, cross-platform binaries with CGO disabled.

## Architecture

### Entry points

- **Library:** root `gomarkdoc` package exposes `Renderer`, its option functions, and rendering methods.
- **CLI:** `cmd/gomarkdoc/main.go` builds and executes one Cobra command; errors reach `log.Fatal` for a non-zero exit.
- **Automation:** `Taskfile.yml` owns build, lint, test, documentation generation, and generated-document verification.

### Data flow

1. Cobra/Viper assemble CLI and configuration-file options in `cmd/gomarkdoc/command.go`.
2. The command expands local and remote package specifiers, excludes requested directories, and resolves an output filename template per package.
3. `go/build` loads each package with selected build tags. `lang.NewPackageFromBuild` parses its selected Go files, creates `go/doc` data, examples, symbols, source positions, and optional repository locations.
4. The CLI constructs a `Renderer` with template and format overrides. The renderer applies named `text/template` templates to the `lang` model.
5. Output is printed to stdout, written to a file, checked against an existing file with a terminal diff, or embedded into Markdown markers.

### Package responsibilities

- `lang/`: domain model and parsing/transformation of Go documentation, source locations, symbols, examples, and repository metadata.
- `format/`: backend-specific Markdown syntax behind `format.Format`; `format/formatcore` holds shared formatting primitives.
- Root package: template registry and renderer orchestration.
- `cmd/gomarkdoc/`: process-level concerns—arguments, configuration, package discovery, filesystem I/O, output checks, and embedding.
- `logger/`: narrow Logrus-backed logger interface.
- `templates/`: embedded `.gotxt` templates; their names form the customization surface.
- `testData/`: fixture Go packages plus Markdown golden files for formatter and CLI behavior.

## Conventions (Observed)

### Error handling

- Functions return errors and wrap contextual I/O, parsing, template, and path errors with `%w` where useful; errors bubble to the CLI boundary.
- Expected absence is handled locally: an absent config file is ignored, an absent embed target starts a new file, and wildcard package paths that load no package are skipped.
- The CLI's `--check` mismatch is intentionally returned after all files are processed so diffs can be printed first.
- No global recovery or structured error type exists; this is a single-process CLI/library, not an HTTP service.

### API and type design

- There is no network API. The public API is Go structs, getters, functional options, and a small formatting interface.
- Exported types and methods use Go doc comments. Internals use concrete types; interfaces are narrow (`logger.Logger`, `format.Format`) and bound at integration points.
- `any` / `interface{}` is limited to template-function registration and Logrus-compatible logging fields, where dynamic values are inherent.
- Options use the conventional `With…` pattern and can return validation errors (`RendererOption`, `PackageOption`, `ConfigOption`).

### I/O and observability

- Filesystem work is concentrated in the CLI; core rendering produces strings and has no process I/O.
- Logging is human-readable, color-capable stderr text—not JSON or metrics/tracing. Per-package CLI logs add a `dir` field.
- The check mode prints a human-readable terminal diff to stderr; normal output can remain clean on stdout.

### Testing and quality

- Tests are co-located (`*_test.go`) and use `github.com/matryer/is` assertions. They exercise parsers, language-model construction, formats, renderer overrides, and CLI flows.
- Golden fixtures live under `testData/`, including expected output for all three Markdown formats, embedding, recursive packages, tags, generics, and unexported symbols.
- `task test` runs `go test -count 1` across all packages with a coverage profile. CI runs Task-based linting, generated-document verification, tests, and Codecov upload on Linux, macOS, and Windows.
- `.golangci.yml` is deliberately strict and enables a broad set of correctness, security, complexity, and style linters.

## Signals / Active Considerations

- **Toolchain drift:** `go.mod` declares Go 1.26, but GitHub workflows still install Go 1.20 and use old action versions. Align CI and release toolchains before relying on newer language or standard-library features.
- **Deprecated API debt:** CLI code uses `io/ioutil`, deprecated since Go 1.16; prefer `os.ReadFile`, `os.WriteFile`, and `os.ReadDir` for touched code.
- **Configuration lifecycle:** Viper is process-global and configuration is read inside `RunE`. This is fine for the one-command executable but makes repeated in-process command execution stateful; reset Viper explicitly if adding more commands or stronger command tests.
- **Package discovery:** Recursive local paths walk directories manually and only ignore `.git`; large repositories or generated/vendor trees may need explicit exclusions or a more complete ignore policy.
- **Rendering extension point:** Adding an output dialect requires a `format.Format` implementation and CLI selection; adding a document component requires a model/template pair and possibly index integration.
- **Known cleanup candidates:** Source comments flag filtering of declarations, Git remote selection, and some rendering edge cases as TODOs. Validate behavior with fixtures before changing parser or template logic.
