// stats.go adds summary statistics helpers. Deliberately shipped WITHOUT tests so
// the pull request's patch coverage is 0% while total coverage barely moves — the
// case a diff-coverage gate is supposed to catch. See ../SECURITY-FIXTURES.md.
package calc

import "sort"

// Sum returns the total of xs.
func Sum(xs []int) int {
	total := 0
	for _, x := range xs {
		total += x
	}
	return total
}

// Mean returns the arithmetic mean of xs, or 0 for an empty slice.
func Mean(xs []int) float64 {
	if len(xs) == 0 {
		return 0
	}
	return float64(Sum(xs)) / float64(len(xs))
}

// Median returns the middle value of xs, averaging the two central values when
// the length is even. Returns 0 for an empty slice. xs is not modified.
func Median(xs []int) float64 {
	if len(xs) == 0 {
		return 0
	}
	sorted := make([]int, len(xs))
	copy(sorted, xs)
	sort.Ints(sorted)

	mid := len(sorted) / 2
	if len(sorted)%2 == 1 {
		return float64(sorted[mid])
	}
	return float64(sorted[mid-1]+sorted[mid]) / 2
}

// MinMax returns the smallest and largest values in xs. The second return value
// is false when xs is empty.
func MinMax(xs []int) (int, int, bool) {
	if len(xs) == 0 {
		return 0, 0, false
	}
	low, high := xs[0], xs[0]
	for _, x := range xs[1:] {
		if x < low {
			low = x
		}
		if x > high {
			high = x
		}
	}
	return low, high, true
}
