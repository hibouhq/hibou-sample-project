# hibou sample project

A small **polyglot fixture** for iterating on [Hibou](https://github.com/hibouhq/hibou)
CI ingestion (coverage, JUnit, SBOMs, SARIF, snapshots) without paying the cost of
dogfooding the real repo. Everything here is deliberately tiny so CI runs in seconds.

## What's inside

| Dir | Language | Test runner | Coverage output |
|-----|----------|-------------|-----------------|
| [`go/`](go/) | Go | `go test` / `gotestsum` | `go/coverage.out`, `go/junit-go.xml` |
| [`ts/`](ts/) | TypeScript | Vitest | `ts/coverage/lcov.info`, `ts/junit-front.xml` |
| [`java/`](java/) | Java | Maven + JUnit 5 | `java/target/site/jacoco/jacoco.xml`, `java/target/surefire-reports/*.xml` |
| [`rust/`](rust/) | Rust | `cargo` + `cargo-llvm-cov` | `rust/lcov.info` |

Each app is the same tiny `calc` library (add/sub/mul/divide/classify) so the four
languages are trivially comparable. **Coverage is intentionally < 100%** in every app
(one branch/function left untested) so coverage diffing has signal.

## ⚠️ Intentional insecurity

Every app also ships an `insecure` module with **deliberately vulnerable code and
dependencies** so security scanners have something to flag. This is a scanner-testing
fixture — see [SECURITY-FIXTURES.md](SECURITY-FIXTURES.md). **Nothing here is real:**
credentials are AWS's public *example* keys (non-functional). Do not copy any of it.

## Quick start

```bash
make test        # run every language's tests
make coverage    # run tests + emit coverage artifacts
make ci          # coverage + SBOM + SARIF (mirrors GitHub Actions)
make clean
```

Missing a toolchain locally? Each `make` target skips languages whose tools aren't
installed (and says so), so you can iterate on one language at a time.

## Release snapshots

The CI workflow also triggers on `release: published` and uploads a Hibou snapshot for
the release ref/SHA **without** attaching any binary artifacts to the GitHub release —
so you can POC snapshot persistence across a release ref cheaply.
