package com.hibou.sample;

import java.io.IOException;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Random;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

/**
 * INTENTIONALLY insecure patterns for scanner testing. Credentials are AWS's
 * public documentation EXAMPLE values (non-functional). Do not copy any of this
 * into production code. See ../SECURITY-FIXTURES.md.
 */
public final class Insecure {

    // Hardcoded credentials — secret scanners flag these.
    public static final String AWS_ACCESS_KEY_ID = "AKIAIOSFODNN7EXAMPLE";
    public static final String AWS_SECRET_ACCESS_KEY = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY";

    // log4j 2.14.1 — Log4Shell (CVE-2021-44228). Logging attacker-controlled input
    // is the exact sink the CVE exploits.
    private static final Logger LOG = LogManager.getLogger(Insecure.class);

    private Insecure() {
    }

    /** Weak hashing (MD5). */
    public static String weakHash(String s) throws NoSuchAlgorithmException {
        MessageDigest md = MessageDigest.getInstance("MD5");
        byte[] digest = md.digest(s.getBytes());
        StringBuilder sb = new StringBuilder();
        for (byte b : digest) {
            sb.append(String.format("%02x", b));
        }
        return sb.toString();
    }

    /** Insecure, non-cryptographic RNG used as a "token". */
    public static int weakToken() {
        return new Random().nextInt();
    }

    /** Command injection — caller-controlled string handed to a shell. */
    public static Process runUserCommand(String userInput) throws IOException {
        return Runtime.getRuntime().exec(userInput);
    }

    /** Logs untrusted input through log4j — the Log4Shell sink. */
    public static void logUserInput(String userInput) {
        LOG.info("user said: {}", userInput);
    }
}
