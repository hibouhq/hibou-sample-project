# Security fixtures — intentional findings

This repo **intentionally** contains vulnerable code, vulnerable dependencies, and
fake secrets, so that Hibou's security ingestion (gosec/trivy SARIF, SBOM CVE
matching, secret scanning) has real findings to work with.

> **None of this is exploitable in the wild.** All credentials are AWS's own
> documented *example* keys (`AKIAIOSFODNN7EXAMPLE`), which are non-functional. All
> "insecure" functions are unused dead code paths kept out of the tested surface.
> This file is the allow-list of what scanners are *expected* to flag.

## Vulnerable dependencies (SBOM / trivy / osv / cargo-audit)

| Language | Package | Version | Advisory |
|----------|---------|---------|----------|
| Go | `github.com/dgrijalva/jwt-go` | v3.2.0 | CVE-2020-26160 (unmaintained; auth bypass) |
| Go | `gopkg.in/yaml.v2` | v2.2.2 | CVE-2019-11253 (billion laughs / DoS) |
| TypeScript | `lodash` | 4.17.11 | CVE-2019-10744 (prototype pollution) |
| Java | `org.apache.logging.log4j:log4j-core` | 2.14.1 | CVE-2021-44228 (Log4Shell) |
| Rust | `time` | 0.1.45 | RUSTSEC-2020-0071 (segfault) |
| Rust | `arrayvec` | 0.4.10 | RUSTSEC-2019-0011 (memory corruption; reachability fixture — never referenced) |

## Reachability fixtures (call-graph analysis)

`hibou analyze reachability` answers "is the vulnerable symbol actually
callable from this code?" — the difference between a CVE that blocks a merge
and one that is noise. Each language carries a **reached / not-reached pair on
real vulnerable dependencies**, so both verdicts are demonstrable:

| Language | File | REACHABLE | NOT reachable |
|---|---|---|---|
| Go (symbol-level) | [`go/reachability.go`](go/reachability.go) | `ParseConfig` → `yaml.Unmarshal` (CVE-2019-11253 and the other yaml.v2 advisories) | `jwt-go` CVE-2020-26160 — the audience-claim path is never called (only signing is) |
| Java (class-level) | [`java/…/Reachability.java`](java/src/main/java/com/hibou/sample/Reachability.java) | `logRequest` → `Logger.info` — the app's bytecode references log4j packages (CVE-2021-44228 marks reachable) | class-level only: a log4j-free module would mark it not reachable |
| Rust (crate-level) | [`rust/src/reachability.rs`](rust/src/reachability.rs) | `time` 0.1.45 (RUSTSEC-2020-0071) — named from source | `arrayvec` 0.4.10 (RUSTSEC-2019-0011) — in the lockfile, never named from source |

Verified with govulncheck 2026-08-06 on the Go module:

```
unreachable GO-2020-0017 [CVE-2020-26160]   ← jwt-go, signing only
REACHABLE   GO-2021-0061 [CVE-2021-4235]    ← yaml.v2 via ParseConfig
REACHABLE   GO-2022-0956 [CVE-2022-3064]    ← yaml.v2 via ParseConfig
```

Each language reports at its honest precision: Go proves the vulnerable
*symbol* on/off the call path (govulncheck); Java proves whether the compiled
app's bytecode *references the packages* of the vulnerable artifact; Rust
proves whether the workspace source *names the crate*. The Java/Rust
reference inventories are stored server-side per snapshot, so a CVE published
AFTER the CI run is classified against unchanged code by the daily advisory
scan — no re-push needed (Go stays symbol-precise and re-verifies on the next
CI run).

Multi-module note: Java scans recurse (point at the repo root — every
module's `target/classes` / `build/classes` and jars are found, test-classes
and unpacked deps skipped); Rust's `cargo metadata` covers all workspace
crates from the root; Go analyzes one module per invocation.

Keep the contracts intact when editing: **do not** call
`jwt.Parser.ParseWithClaims` or `MapClaims.VerifyAudience` in the Go module,
and **do not** `use arrayvec` in the Rust crate — or the unreachable fixtures
flip.

## SAST findings (gosec / semgrep / trivy / clippy)

Present in every `insecure.*` module:

- **Hardcoded credentials** — AWS example access key + secret (secret scanners).
- **Weak hashing** — MD5 used for hashing.
- **Insecure RNG** — non-cryptographic random used for a token.
- **Command injection** — user input passed to a shell.
- **Code injection** (TS only) — `eval()` of caller-supplied string.
- **Vulnerable-dep usage** — the CVE deps above are actually called.

## What is NOT a finding

- The `calc` libraries are clean.
- Coverage < 100% is intentional (untested `classify(negative)` branch + the
  `insecure` module), not a bug.
