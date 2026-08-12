package com.hibou.sample;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.apache.logging.log4j.util.Strings;

/**
 * Call-graph fixture: pairs a REACHED vulnerable sink with an UNREACHED one on
 * the same vulnerable dependency (log4j-core 2.14.1, CVE-2021-44228), so a Java
 * reachability analyzer has both verdicts to report.
 *
 * <p>Nothing here is exploitable — inputs are literals and no attacker-controlled
 * data reaches a sink at runtime. See ../SECURITY-FIXTURES.md.
 */
public final class Reachability {

    private static final Logger LOG = LogManager.getLogger(Reachability.class);

    private Reachability() {
    }

    // ── REACHABLE ────────────────────────────────────────────────────────
    /**
     * Calls {@code Logger.info(String)} with an interpolated value — the exact
     * sink Log4Shell (CVE-2021-44228) exploits via JNDI lookup in message text.
     * A call-graph analyzer traces this from a public method to the vulnerable
     * code.
     *
     * <p>Expected Hibou verdict: reachability = "reachable" for CVE-2021-44228.
     */
    public static void logRequest(String requestId) {
        LOG.info("handling request {}", requestId);
    }

    // ── UNREACHABLE ──────────────────────────────────────────────────────
    /**
     * Touches log4j's {@code Strings} utility only. The class comes from
     * log4j-api, and nothing on this path reaches the message-lookup code where
     * CVE-2021-44228 lives — the dependency is matched by the SBOM, but the
     * vulnerable symbol is never called from here.
     *
     * <p>Expected Hibou verdict: reachability = "not_reachable" for any advisory
     * whose vulnerable symbol is the lookup path, when this is the ONLY log4j
     * usage in the analyzed module. (In this fixture {@link #logRequest} keeps
     * the overall verdict reachable on purpose — the two methods exist so both
     * outcomes are demonstrable; analyze them separately to see the contrast.)
     */
    public static boolean isBlank(String value) {
        return Strings.isBlank(value);
    }
}
