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
| Java | `org.apache.logging.log4j:log4j-core` | 2.14.1 | CVE-2021-44228 (Log4Shell) |
| Rust | `time` | 0.1.45 | RUSTSEC-2020-0071 (segfault) |

Two fixtures were remediated to exercise the "vulnerability resolved" path:
`gopkg.in/yaml.v2` v2.2.2 → v2.4.0 (CVE-2019-11253) and `lodash` 4.17.11 →
4.17.21 (CVE-2019-10744). Both are still imported, so the dependency stays in
the SBOM at a non-vulnerable version rather than disappearing from it.

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
