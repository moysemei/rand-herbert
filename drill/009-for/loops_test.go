package loops

import "testing"

func TestSumTo(t *testing.T) {
	if got := SumTo(5); got != 15 {
		t.Errorf("SumTo(5) = %d, want 15", got)
	}
	if got := SumTo(1); got != 1 {
		t.Errorf("SumTo(1) = %d, want 1", got)
	}
	if got := SumTo(0); got != 0 {
		t.Errorf("SumTo(0) = %d, want 0", got)
	}
	if got := SumTo(100); got != 5050 {
		t.Errorf("SumTo(100) = %d, want 5050", got)
	}
}

func TestFactorial(t *testing.T) {
	if got := Factorial(5); got != 120 {
		t.Errorf("Factorial(5) = %d, want 120", got)
	}
	if got := Factorial(1); got != 1 {
		t.Errorf("Factorial(1) = %d, want 1", got)
	}
	if got := Factorial(0); got != 1 {
		t.Errorf("Factorial(0) = %d, want 1", got)
	}
	if got := Factorial(10); got != 3628800 {
		t.Errorf("Factorial(10) = %d, want 3628800", got)
	}
}

func TestSumOfMultiples(t *testing.T) {
	if got := SumOfMultiples(10); got != 23 {
		t.Errorf("SumOfMultiples(10) = %d, want 23", got)
	}
	if got := SumOfMultiples(16); got != 60 {
		t.Errorf("SumOfMultiples(16) = %d, want 60", got)
	}
	if got := SumOfMultiples(1); got != 0 {
		t.Errorf("SumOfMultiples(1) = %d, want 0", got)
	}
	if got := SumOfMultiples(1000); got != 233168 {
		t.Errorf("SumOfMultiples(1000) = %d, want 233168", got)
	}
}
