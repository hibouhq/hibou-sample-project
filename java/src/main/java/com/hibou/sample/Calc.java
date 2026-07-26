package com.hibou.sample;

/** A tiny arithmetic library, mirrored across every language in this fixture. */
public final class Calc {

    private Calc() {
    }

    public static int add(int a, int b) {
        return a + b;
    }

    public static int sub(int a, int b) {
        return a - b;
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
    public static String classify(int n) {
        if (n < 0) {
            return "negative";
        }
        if (n == 0) {
            return "zero";
        }
        return "positive";
    }
}
