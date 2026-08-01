package com.hibou.sample;

/** A tiny arithmetic library, mirrored across every language in this fixture. */
public final class Calc {

    private Calc() {
    }

    public static int add(int a, int b) {
        return a + b;
    }

    public static int mul(int a, int b) {
        return a * b;
    }

    /** Returns a / b, throwing ArithmeticException when b == 0. */
    public static int divide(int a, int b) {
        if (b == 0) {
            throw new ArithmeticException("divide by zero");
        }
        return a / b;
    }

    /**
     * Labels n as "negative", "zero", or "positive". The "negative" branch is
     * intentionally left uncovered by tests so coverage stays below 100%.
     */
    /**
     * Clamps n into [lo, hi]. The compound condition sits on one line and the
     * tests exercise only some of its outcomes, so JaCoCo reports the line as
     * PARTIALLY covered (mb > 0 && cb > 0) — the fixture for partial-branch
     * rendering.
     */
    public static int clamp(int n, int lo, int hi) {
        if (n < lo || n > hi) {
            return n < lo ? lo : hi;
        }
        return n;
    }

    /**
     * Partial-branch fixtures: each method keeps several branch outcomes on
     * one line while tests exercise only a subset, so JaCoCo reports the line
     * with mb &gt; 0 and cb &gt; 0 (partially covered).
     */
    public static boolean isValidPort(int n, boolean allowPrivileged) {
        return n > 0 && n < 65536 && (allowPrivileged || n >= 1024);
    }

    public static String grade(int score) {
        return score >= 90 ? "A" : score >= 75 ? "B" : score >= 50 ? "C" : "F";
    }

    public static String classify(int n) {
        if (n < 0) {
            return "negative";
        }
        if (n == 0) {
            return "zero";
        }
        if (n == 1) {
            return "one";
        }
        return "positive";
    }

    /** New in this PR, covered by a test — renders as "new covered". */
    public static int pow(int base, int exp) {
        int result = 1;
        for (int i = 0; i < exp; i++) {
            result *= base;
        }
        return result;
    }

    /** New in this PR and INTENTIONALLY untested — renders as "new uncovered". */
    public static boolean isEven(int n) {
        return n % 2 == 0;
    }

    /**
     * New in this PR — textbook PARTIAL coverage: tests only pass n == 0, so the
     * first operand is always true and `n == 999` is NEVER reached (the ||
     * short-circuits). The line runs but the second branch outcome is missing.
     */
    public static boolean isSentinel(int n) {
        if (n == 0 || n == 999) {
            return true;
        }
        return false;
    }
}
