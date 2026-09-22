// Package calc is a tiny arithmetic library used to exercise coverage tooling.
package calc

import "errors"

// ErrDivideByZero is returned by Divide when the divisor is zero.
var ErrDivideByZero = errors.New("divide by zero")

// Add returns a + b.
func Add(a, b int) int { return a + b }

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
	case n == 1:
		return "one"
	default:
		return "positive"
	}
}

// Pow returns base**exp for exp >= 0. New in this PR and covered by a test, so
// its lines render as "new covered" in the comparison.
func Pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

// IsEven reports whether n is even. New in this PR and INTENTIONALLY left
// untested, so its lines render as "new uncovered" — the regression to fix.
func IsEven(n int) bool {
	return n%2 == 0
}

// IsSentinel reports whether n is a sentinel value. The test only ever passes 0,
// so `n == 0` is always true and `|| n == 999` is NEVER reached. NOTE: Go's
// coverage is line-level (no branch coverage), so this renders as covered — its
// TS and Java twins show the same line as PARTIAL.
func IsSentinel(n int) bool {
	if n == 0 || n == 999 {
		return true
	}
	return false
}
