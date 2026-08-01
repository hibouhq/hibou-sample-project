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

## OS package vulnerabilities (image SBOMs)

`images/{alpine,debian,rhel}/Dockerfile` build INTENTIONALLY old base images so
their OS packages carry known CVEs in the three distro ecosystems. CI emits one
CycloneDX SBOM per image (`<os>-image-sbom.cdx.json`); each OS package's purl
carries a `distro=` qualifier so Hibou matches advisories **release-scoped** and
reconciles OSV + Trivy sources onto one advisory.

| Image | Ecosystem | Release scope | Notes |
|-------|-----------|---------------|-------|
| `alpine:3.9` | apk | Alpine 3.9 | old `openssl`/`libssl`; also installs a vulnerable `dockerize` v0.6.1 binary from GitHub releases |
| `debian:10` | deb | Debian 10 | EOL buster `openssl`/`libc` packages |
| `redhat/ubi8:8.5` | rpm | RHEL 8 | old `openssl-libs`/`glibc`/`platform-python` rpms |

The release scoping is the point: an Alpine 3.9 fix version must not match a
package from Alpine 3.18. `dockerize v0.6.1` additionally exercises Go-binary
detection (bundled `golang.org/x/*` advisories) alongside the OS packages.

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
