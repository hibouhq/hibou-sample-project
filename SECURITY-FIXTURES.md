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
| TypeScript | `jquery` | 3.4.1 | CVE-2020-11023 (XSS; **known-exploited**, see below) |
| Java | `org.apache.logging.log4j:log4j-core` | 2.14.1 | CVE-2021-44228 (Log4Shell; **known-exploited**, see below) |
| Rust | `time` | 0.1.45 | RUSTSEC-2020-0071 (segfault) |

## Known-exploited vulnerabilities (KEV)

Two of the dependencies above are on the public exploited-vulnerabilities
catalogs (CISA KEV; the ENISA EUVD exploited set lists them too), which Hibou
mirrors and gates on unconditionally — a new finding on a listed advisory fails
the PR gate regardless of severity thresholds, and the finding detail carries
an "exploited" badge with the listing facts.

| Package | CVE | CVSS severity | Why it is the fixture |
|---------|-----|---------------|----------------------|
| `jquery` 3.4.1 (ts) | CVE-2020-11023 | **medium** (6.1) | The interesting case: passes every severity-only gate (`max_critical`/`max_high`) and must STILL block, because exploitation in the wild is documented. If this one gets through, the exploited block is broken. |
| `log4j-core` 2.14.1 (java) | CVE-2021-44228 | critical (10.0) | Also carries CISA's known-ransomware-campaign flag and a BOD 22-01 due date, so the detail view's richest KEV facts render. Blocks via severity gates anyway — the KEV reason should appear alongside, not instead. |

`jquery` is declared in `ts/package.json` but deliberately **not imported**:
jQuery needs a DOM and would crash this Node-only fixture, and dependency
matching reads the SBOM/lockfile anyway. Waiving either finding through the
acceptance flow is the sanctioned way past the block and should surface the
exploited warning in the dialog.

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
