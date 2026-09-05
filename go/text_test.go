package calc

import "testing"

func TestReverse(t *testing.T) {
	cases := map[string]string{
		"":      "",
		"a":     "a",
		"abc":   "cba",
		"école": "elocé",
	}
	for in, want := range cases {
		if got := Reverse(in); got != want {
			t.Errorf("Reverse(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := map[string]bool{
		"":                               true,
		"racecar":                        true,
		"A man, a plan, a canal: Panama": true,
		"Was it a car or a cat I saw?":   true,
		"12 21":                          true,
		"hello":                          false,
		"palindrome":                     false,
	}
	for in, want := range cases {
		if got := IsPalindrome(in); got != want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", in, got, want)
		}
	}
}

func TestWordCount(t *testing.T) {
	cases := map[string]int{
		"":                 0,
		"   ":              0,
		"one":              1,
		"one two three":    3,
		"  padded  words ": 2,
	}
	for in, want := range cases {
		if got := WordCount(in); got != want {
			t.Errorf("WordCount(%q) = %d, want %d", in, got, want)
		}
	}
}
