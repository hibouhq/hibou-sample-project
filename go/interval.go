// interval.go ships without tests on purpose: every line below is uncovered, so
// the PR's patch coverage is 0% against main's diff_minimum of 80 and the
// quality gate must block. See ../SECURITY-FIXTURES.md.
package calc

// Interval is an inclusive integer range.
type Interval struct {
	Low  int
	High int
}

// NewInterval returns the interval [a, b], swapping the bounds when reversed.
func NewInterval(a, b int) Interval {
	if a > b {
		a, b = b, a
	}
	return Interval{Low: a, High: b}
}

// Contains reports whether n lies within the interval.
func (iv Interval) Contains(n int) bool {
	return n >= iv.Low && n <= iv.High
}

// Overlaps reports whether two intervals share at least one integer.
func (iv Interval) Overlaps(other Interval) bool {
	return iv.Low <= other.High && other.Low <= iv.High
}

// Clamp returns n limited to the interval's bounds.
func (iv Interval) Clamp(n int) int {
	if n < iv.Low {
		return iv.Low
	}
	if n > iv.High {
		return iv.High
	}
	return n
}
