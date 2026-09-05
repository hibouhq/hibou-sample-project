// text.go adds string helpers, fully covered by text_test.go. The counterpart to
// stats.go: same shape of change, but every new line is exercised, so patch
// coverage is 100% and a diff-coverage gate should pass.
package calc

import "strings"

// Reverse returns s with its runes in reverse order.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome reports whether s reads the same in both directions, ignoring
// case and any character that is not a letter or digit.
func IsPalindrome(s string) bool {
	var cleaned []rune
	for _, r := range strings.ToLower(s) {
		if isAlphanumeric(r) {
			cleaned = append(cleaned, r)
		}
	}
	for i, j := 0, len(cleaned)-1; i < j; i, j = i+1, j-1 {
		if cleaned[i] != cleaned[j] {
			return false
		}
	}
	return true
}

// WordCount returns the number of whitespace-separated words in s.
func WordCount(s string) int {
	return len(strings.Fields(s))
}

func isAlphanumeric(r rune) bool {
	switch {
	case r >= 'a' && r <= 'z':
		return true
	case r >= '0' && r <= '9':
		return true
	default:
		return false
	}
}
