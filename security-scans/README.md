# Multilanguage scanner fixtures

These are intentionally unsafe source snippets for static analysis only. Never
execute/import their `inspect`, `execute`, or command methods. The harness only
copies files and invokes a scanner. No fixture is part of an application build.

## Run

```sh
make sast-fixtures-test
make sast-fixtures OPENGREP=/absolute/path/to/opengrep
make sast-fixtures-interfile OPENGREP=/absolute/path/to/interfile-enabled/opengrep
```

Python 3 and the selected scanner executable are required. Unlike optional demo
targets, a missing scanner fails this test suite. No registry rules are fetched:
all rules are local. The suite does not install/download tools. Acquire a pinned,
verified binary ahead of time, including from an approved internal mirror or
preseeded offline location, and pass its absolute path through `OPENGREP`.

Every case runs in an isolated temporary project root to avoid class-name/import
collisions with other fixtures. The harness asserts exact rule/file finding
multisets, rejects engine diagnostics and nonzero exits, and checks that every
fixture source was scanned. Unsupported modes are recorded as skipped, not
passed. `sast-fixtures-interfile` fails immediately if the binary does not expose
`--taint-interfile`. Machine-readable output is `results.json` (ignored by Git);
stale output is removed at the start of each invocation.

## Matrix

30 source files: 14 Java, 8 Python, 8 JavaScript. There are 24 isolated cases,
each checked in basic, intrafile, and interfile modes when advertised by engine.

| Case | Basic | Intrafile | Interfile |
|---|---|---|---|
| Direct source to command sink | Finding | Finding | Finding |
| Branch joins source and constant | Finding | Finding | Finding |
| Loop-assigned input (Java) | Finding | Finding | Finding |
| Source exists but sink gets constant | No finding | No finding | No finding |
| Explicitly modeled sanitizer returns constant | No finding | No finding | No finding |
| Input reassigned to constant | No finding | No finding | No finding |
| Same-file helper receives source | No finding | Finding | Finding |
| Cross-file helper receives source | No finding | No finding | Finding |
| Cross-file return (Java) | No finding | No finding | Finding |
| Cross-file inherited method (Java) | No finding | No finding | Finding |

Rules intentionally model only one environment variable as source and command
execution as sink. Sanitizer behavior is explicitly declared in rules; this does
not prove recognition of arbitrary real validators/frameworks. Expected misses
in weaker modes are capability controls, not comprehensive vulnerability coverage.

## Executed results (2026-10-09)

| Engine | Executed checks | Skipped checks | Result |
|---|---|---|---|
| OpenGrep 1.30.2 Linux x86 release | 48 (basic + intrafile) | 24 (interfile) | All executed checks passed |
| Semgrep 1.172.0 installed CLI | 24 (basic) | 48 (intrafile + interfile flags absent) | All executed checks passed |

OpenGrep binary was acquired from the upstream `v1.30.2` release and verified
against published SHA-256:
`a66aa3278457f02b287b985a45b6762aebcaba5000f2689245fd1ed86d1456c7`.
This is checksum verification, not a claim that Cosign provenance was verified.

The release does not expose `--taint-interfile`; the adjacent development checkout
does. No cross-file positive detection has been executed here. Tests are ready
for a compatible build; do not relabel these skips as successful interfile tests.

A separate source scan of the sample Java application with intrafile taint
enabled still detected the ordinary weak-hash pattern and emitted native SARIF
2.1.0. No JDK or `.class` files were used. SpotBugs requires compiled bytecode
and has not been run in this container; Hibou ingestion remains unverified.

12 standard-library harness unit tests cover expectation modes, negative controls,
error rejection, duplicates, missing findings, rule IDs, isolation and incomplete
source coverage. This synthetic matrix is not a large-repository performance
benchmark, independent security corpus, framework coverage study, or parity claim.