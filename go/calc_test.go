package calc

import "testing"

func TestAdd(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Fatalf("Add(2,3) = %d, want 5", Add(2, 3))
	}
}

func TestSub(t *testing.T) {
	if Sub(5, 3) != 2 {
		t.Fatalf("Sub(5,3) = %d, want 2", Sub(5, 3))
	}
}

func TestMul(t *testing.T) {
	if Mul(4, 3) != 12 {
		t.Fatalf("Mul(4,3) = %d, want 12", Mul(4, 3))
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
	// Classify(negative) is intentionally NOT tested → keeps coverage < 100%.
}

func TestWeakHash(t *testing.T) {
	if WeakHash("x") == "" {
		t.Fatal("WeakHash returned empty string")
	}
}
