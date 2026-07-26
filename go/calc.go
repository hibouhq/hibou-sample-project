// Package calc is a tiny arithmetic library used to exercise coverage tooling.
package calc

import "errors"

// ErrDivideByZero is returned by Divide when the divisor is zero.
var ErrDivideByZero = errors.New("divide by zero")

// Add returns a + b.
func Add(a, b int) int { return a + b }

// Sub returns a - b.
func Sub(a, b int) int { return a - b }

// Mul returns a * b.
func Mul(a, b int) int { return a * b }

// Divide returns a / b, or ErrDivideByZero when b == 0.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// Classify labels n as "negative", "zero", or "positive". The "negative" branch
// is intentionally left uncovered by tests so the fixture reports < 100% coverage.
func Classify(n int) string {
	switch {
	case n < 0:
		return "negative"
	case n == 0:
		return "zero"
	default:
		return "positive"
	}
}
