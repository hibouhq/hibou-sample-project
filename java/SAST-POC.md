# Java SAST POC

Shortlist: OpenGrep for multi-language source analysis, and SpotBugs + Find
Security Bugs for complementary JVM bytecode analysis. Dependency analysis stays
with the existing Trivy/SBOM pipeline; SAST is not a Log4Shell dependency detector.

## Run

Use JDK 17 and Maven for bytecode analysis. The Maven profile pins its plugin and
Find Security Bugs versions; ordinary tests do not enable this profile.

```sh
make java-spotbugs
make java-opengrep OPENGREP=/absolute/path/to/interfile-enabled/opengrep
make java-sast OPENGREP=/absolute/path/to/interfile-enabled/opengrep
make -j2 java-sast OPENGREP=/absolute/path/to/interfile-enabled/opengrep
```

The parallel invocation overlaps source scanning with compile/bytecode analysis;
the latter still runs compilation before SpotBugs. Do not run `java-sast` alongside
`java-cov` in the same workspace: both Maven invocations can mutate `java/target`.
This POC compiles for its own bytecode scan; production integration should reuse
existing compatible build artifacts as proposed in ADR-035.

OpenGrep must expose `scan --taint-interfile`. The adjacent source checkout
contained this flag at `f759d38bd` (`v1.30.0-326-gf759d38bd`); this is evidence
about that development checkout, not a claim about released 1.30.x binaries.
Check your binary's help before running. Do not silently fall back to intrafile.

Outputs at repository root: `java-opengrep.sarif`, `java-spotbugs.sarif`.
Existing CI upload globs accept these reports when present. CI does not yet
install either scanner; these are local opt-in POC targets. `make sarif` invokes
OpenGrep if installed; `make java-sast` additionally invokes bytecode analysis.
Missing tools print a skip, not evidence of a successful scan. Scanner execution
errors propagate. Delete old reports before evaluation to avoid stale evidence.

## Expected checks

| Case | OpenGrep expectation | Purpose |
|---|---|---|
| `CommandController.handle` reads `request.getParameter("command")` and calls `CommandExecutor.run` | Command-injection finding at `Runtime.exec` in another file | Cross-file source-to-sink flow |
| Same command rule with `--disable-interfile --disable-intrafile` | No command-injection finding | Distinguish flow analysis from sink matching |
| Replace controller argument with constant `"echo fixture"` in a disposable copy | No command-injection finding | Negative control |
| `CommandController.handleConstant` | No additional source-derived flow from this entry point | Constant-input control (shared sink still reports vulnerable entry) |
| Existing `Insecure.weakHash` | Weak-hash finding | Ordinary source pattern |
| Existing `Calc` | No findings from these two rules | Clean control |

The request class is a minimal fixture, not a servlet implementation. The source
rule intentionally models `getParameter`; it does not establish framework-wide
coverage. All command methods are unused, never invoked by tests. Scan them;
do not execute them.

SpotBugs + Find Security Bugs should be inspected for weak hashing, weak random
tokens, and command-execution warnings. Record what it actually reports rather
than assuming it recognizes this custom request class as an untrusted source.

For each scanner record engine/plugin/ruleset versions, elapsed time, errors,
skipped files, expected detections, unexpected detections, SARIF source paths,
fingerprints, and dataflow traces when available. Upload to a development Hibou
instance and verify source locations and producer identity before claiming
ingestion compatibility. OpenGrep 1.30.2 source tests and native SARIF output have
now executed successfully; that release supports intrafile but not the interfile
flag required by `java-opengrep`. Cross-file detection and bytecode scanning are
now measured (see Results); ingestion was checked against Hibou's SARIF parser
with these captures, not yet through an upload to a running instance. See
[../security-scans/README.md](../security-scans/README.md) for the measured
multilanguage matrix.

## Results (2026-10-10)

Binaries were downloaded from the upstream release pages and checked against
the published sha256 digests. SpotBugs ran in the `maven:3.9.9-eclipse-temurin-17`
image with `mvn -Psast compile spotbugs:spotbugs`.

| Scanner | Version | Cross-file command injection | Other findings | Fingerprints | Code flows |
|---|---|---|---|---|---|
| OpenGrep release | 1.30.1 | not found (no `--taint-interfile`; `--taint-intrafile` only) | weak hash | `matchBasedId/v1` | — |
| OpenGrep interfile pre-release | `v2.0.0-nopython-interfile.alpha.3` | found at `CommandExecutor.java:9` | weak hash | `matchBasedId/v1` | with `--dataflow-traces`: `CommandController.java:6` → `CommandExecutor.java:9` |
| SpotBugs + Find Security Bugs | 4.9.3 + 1.13.0 | `COMMAND_INJECTION` at `CommandExecutor.java:9` and `Insecure.java:46` | MD5, predictable random, random used once, default encoding | none | — |

Observations that affect integration:

- Interfile taint ships only in the separate `v2.0.0-nopython-interfile`
  pre-release line, which reports its own version as `1.30.0`. A version string
  cannot tell it from the 1.30.0 release; pin binaries by digest.
- Without `--dataflow-traces` the interfile build reports the sink only, with no
  `codeFlows`. Pass the flag: the source file of a flow is needed to decide
  whether a later change repaired it.
- OpenGrep prefixes rule ids with the rules file's directory
  (`java.hibou.java.command-injection`); moving the rules file renames every
  rule and, through `matchBasedId`, every finding identity.
- SpotBugs writes paths relative to the Java source root
  (`com/hibou/sample/CommandExecutor.java`), not to the repository, and no
  fingerprints. `hibou upload` now rewrites such paths to repo-relative ones
  when exactly one checkout file matches, then adds content fingerprints.
- Both scanners report `executionSuccessful: true`; SpotBugs sets `exitCode: 1`
  ("BUGS FOUND"), which is not a failure.

## Comparison boundary

Cross-file capability makes OpenGrep worth evaluating for Java security analysis,
but does not establish GitLab parity. These authored smoke-test rules are not a
production ruleset or an unbiased benchmark. Next evaluation needs a licensed,
pinned ruleset and an independent corpus covering servlet/Spring inputs, SQL,
XSS, path traversal, SSRF, sanitizers, interfaces, overloads, and long call chains.
Measure recall, false positives, runtime, and explainability against identical
commits and configurations. General Java correctness and maintainability require
separate rule coverage; source security taint alone is not that whole capability.