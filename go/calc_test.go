package calc

import "testing"

func TestAdd(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Fatalf("Add(2,3) = %d, want 5", Add(2, 3))
	}
}

func TestDivide(t *testing.T) {
	q, err := Divide(10, 2)
	if err != nil || q != 5 {
		t.Fatalf("Divide(10,2) = %d, %v, want 5, nil", q, err)
	}
}

func TestDivideByZero(t *testing.T) {
	if _, err := Divide(1, 0); err == nil {
		t.Fatal("Divide(1,0) = nil error, want ErrDivideByZero")
	}
}

func TestClassify(t *testing.T) {
	if got := Classify(0); got != "zero" {
		t.Fatalf("Classify(0) = %q, want zero", got)
	}
	if got := Classify(5); got != "positive" {
		t.Fatalf("Classify(5) = %q, want positive", got)
	}
	// Newly covered in this PR: the negative branch was uncovered on main, so it
	// renders as "gained coverage" in the comparison.
	if got := Classify(-5); got != "negative" {
		t.Fatalf("Classify(-5) = %q, want negative", got)
	}
	// Covers the new "one" case added in this PR — "new covered".
	if got := Classify(1); got != "one" {
		t.Fatalf("Classify(1) = %q, want one", got)
	}
}

func TestPow(t *testing.T) {
	if Pow(2, 10) != 1024 {
		t.Fatalf("Pow(2,10) = %d, want 1024", Pow(2, 10))
	}
}

func TestIsSentinel(t *testing.T) {
	// Only n == 0 is passed, so `n == 999` is never reached in `if n == 0 || n == 999`.
	if !IsSentinel(0) {
		t.Fatal("IsSentinel(0) = false, want true")
	}
}

func TestWeakHash(t *testing.T) {
	if WeakHash("x") == "" {
		t.Fatal("WeakHash returned empty string")
	}
}
