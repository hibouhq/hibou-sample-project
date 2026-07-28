package com.hibou.sample;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

import org.junit.jupiter.api.Test;

class CalcTest {

    @Test
    void clampsInRangeAndBelow() {
        // Deliberately no above-range case: "n > hi" never goes true, leaving
        // the compound condition partially covered.
        assertEquals(5, Calc.clamp(5, 0, 10));
        assertEquals(0, Calc.clamp(-3, 0, 10));
    }

    @Test
    void rejectsNonPositivePort() {
        // Only the first condition's false outcome is exercised; the rest of
        // the chain stays partially covered.
        assertEquals(false, Calc.isValidPort(0, false));
    }

    @Test
    void gradesAnA() {
        assertEquals("A", Calc.grade(95));
    }

    @Test
    void adds() {
        assertEquals(5, Calc.add(2, 3));
    }

    @Test
    void subtracts() {
        assertEquals(2, Calc.sub(5, 3));
    }

    @Test
    void multiplies() {
        assertEquals(12, Calc.mul(4, 3));
    }

    @Test
    void divides() {
        assertEquals(5, Calc.divide(10, 2));
    }

    @Test
    void throwsOnDivideByZero() {
        assertThrows(ArithmeticException.class, () -> Calc.divide(1, 0));
    }

    @Test
    void classifies() {
        assertEquals("zero", Calc.classify(0));
        assertEquals("positive", Calc.classify(5));
        // classify(negative) intentionally untested → keeps coverage < 100%.
    }
}
