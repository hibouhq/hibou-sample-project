package com.hibou.sample;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

import org.junit.jupiter.api.Test;

class CalcTest {

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
