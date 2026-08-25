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

## Application software installed into an image (Tomcat)

`images/tomcat/Dockerfile` covers the case the three OS images do not: software
installed into an image that is **not** an OS package and **not** a declared
dependency. Tomcat ships as jars under `/usr/local/tomcat/lib`, so `trivy fs`
over the source tree never sees it — only an image SBOM does, where Trivy's jar
analyzer emits `pkg:maven/org.apache.tomcat/*` components (35 of them).

CI builds the Dockerfile twice, and the pair is the fixture:

| Build | `TOMCAT_TAG` | SBOM | Expected |
|-------|--------------|------|----------|
| `hibou-fixture-tomcat` | `11.0.22-jre21-temurin` (default) | `tomcat-image-sbom.cdx.json` | **none** — a recent release whose 2026 CVEs have no maven coordinates in GHSA/OSV (see below), only an NVD CPE |
| `hibou-fixture-tomcat-old` | `11.0.2-jre21-temurin` | `tomcat-old-image-sbom.cdx.json` | **findings** — OSV/GHSA carry ~20 advisories for `org.apache.tomcat:tomcat-catalina` at 11.0.2, incl. CVE-2025-24813 (RCE, fixed in 11.0.3) |

The default is 11.0.22 because a recent, apparently-clean release is the
interesting case; 11.0.2 exists only to make that empty result interpretable. Alone, an empty 11.0.22 is
ambiguous — missing upstream data and a broken purl match look identical. With
both, findings on 11.0.2 prove the maven image path works end-to-end, so an
empty 11.0.22 isolates the cause to the advisory data.

**Why 11.0.22 is empty (verified 2026-08-24).** It is not lag. The 2026 Tomcat
CVEs that fix in 11.0.23/11.0.24 (CVE-2026-53404, -53434, -55276, -59083,
-59084) exist in every database, but *without package coordinates*:

- GHSA has them (`GHSA-cwxf-7cw2-7r32`, `GHSA-mqg3-r7h5-24x4`,
  `GHSA-4x29-79gh-6v8q`), published 2026-06-29 / 2026-07-14, `severity: critical`
  — with `type: unreviewed` and an **empty `vulnerabilities` array**: no
  ecosystem, no package name, no version range.
- OSV has them by CVE ID, but each `affected` entry carries only `type: GIT`
  commit ranges from Apache's CNA record — no `package`, no `ecosystem: Maven`.

So an OSV *version query* (`org.apache.tomcat:tomcat-catalina` @ `11.0.22`)
returns zero, while a query by CVE ID returns the advisory. NVD does carry the
CPE binding (`cpe:2.3:a:apache:tomcat`, `versionEndExcluding: 11.0.23`), which is
why CPE-based scanners report these and every purl-based scanner — Trivy, Snyk,
Dependabot, Hibou — structurally cannot. The last *reviewed* Tomcat GHSAs are
dated 2026-05-12, and their ranges end at `< 11.0.22`.

> Corollary: `trivy image --scanners vuln` reports **zero** Tomcat CVEs on both
> builds. On 11.0.22 that is the coordinate gap above; on 11.0.2 it is Trivy's
> Java DB. Hibou's own OSV matching over the SBOM purls is what must produce the
> 11.0.2 findings — that is what this fixture proves. Closing the 11.0.22 case
> needs CPE matching, not a better purl.

### How other scanners behave on this fixture

Useful when comparing Hibou against another tool on the same image: on the
11.0.22 build, **a clean result is the norm, not a distinguishing feature.**
Verified 2026-08-24.

| Tool | Result on 11.0.22 | Why |
|---|---|---|
| **Trivy** (`image --scanners vuln`) | none | purl-only matching; its Java DB is GHSA-derived, and the GHSA entries carry no coordinates |
| **Dependency-Track** | none *from a Trivy or Syft CycloneDX SBOM* | its internal analyzer does match CPEs against a mirrored NVD, but ["matching against data from the NVD requires components to have a valid CPE"](https://dependencytrack.github.io/docs/next/reference/analyzers/) — and DT does not infer one from a purl. Feed it an SBOM with `cpe:2.3:a:apache:tomcat:11.0.22` added by hand and it reports all five |
| **Grype** | may report, depending on config | consumes Syft's CPE *candidate set*, which includes the matching `apache:tomcat`. Note Grype ships a per-ecosystem kill switch (`match.java.using-cpes`) because CPE matching is a known false-positive source |
| **Snyk / Dependabot** | none | purl-only |
| **CPE-based scanners** (NVD-backed CSPM/agentless) | **all five** | they match `cpe:2.3:a:apache:tomcat` with `versionEndIncluding 11.0.23` straight from NVD |

Two traps this fixture exists to expose, both worth knowing before trusting any
tool's output here:

1. **Trivy's CycloneDX carries no `cpe` field at all** — 0 of 177 components. A
   pipeline that "ingests CPEs" from a Trivy SBOM is ingesting nothing.
2. **Syft's CycloneDX carries one CPE, and picks a wrong one.** Syft's native
   output gives `tomcat-coyote` three candidates including the matching
   `cpe:2.3:a:apache:tomcat:11.0.22`; its CycloneDX keeps only
   `apache:tomcat-catalina`, which is not in NVD's dictionary. Other jars fare
   worse (`apache:tomcat-el-api:6.0` — the spec version, not the release;
   `apache:catalina-tribes` — an invented product). A CPE-aware pipeline fed
   Syft's CycloneDX still matches nothing, while *looking* correctly configured.

A third trap, if you do wire CPE matching up: `cpe:2.3:a:apache:tomcat` appears
as a candidate on **18 of the 180 artifacts** Syft finds in this image, including
`tomcat-i18n-cs` (a Czech resource bundle). Five CVEs then become 90 findings
unless they are collapsed per product. Dependency-Track hit the identical shape
with Telerik UI for WPF — per-submodule purls against one monolithic product CPE
([discussion #4180](https://github.com/DependencyTrack/dependency-track/discussions/4180)).

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
