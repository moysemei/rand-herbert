package practice

import "testing"

func TestSumEven(t *testing.T) {
	if got := SumEven(6); got != 12 {
		t.Errorf("SumEven(6) = %d, want 12", got)
	}
	if got := SumEven(7); got != 12 {
		t.Errorf("SumEven(7) = %d, want 12", got)
	}
	if got := SumEven(1); got != 0 {
		t.Errorf("SumEven(1) = %d, want 0", got)
	}
	if got := SumEven(10); got != 30 {
		t.Errorf("SumEven(10) = %d, want 30", got)
	}
}

func TestCountMultiples(t *testing.T) {
	if got := CountMultiples(10, 3); got != 3 {
		t.Errorf("CountMultiples(10, 3) = %d, want 3", got)
	}
	if got := CountMultiples(10, 5); got != 2 {
		t.Errorf("CountMultiples(10, 5) = %d, want 2", got)
	}
	if got := CountMultiples(20, 4); got != 5 {
		t.Errorf("CountMultiples(20, 4) = %d, want 5", got)
	}
	if got := CountMultiples(2, 3); got != 0 {
		t.Errorf("CountMultiples(2, 3) = %d, want 0", got)
	}
}

func TestPower(t *testing.T) {
	if got := Power(2, 10); got != 1024 {
		t.Errorf("Power(2, 10) = %d, want 1024", got)
	}
	if got := Power(3, 4); got != 81 {
		t.Errorf("Power(3, 4) = %d, want 81", got)
	}
	if got := Power(7, 1); got != 7 {
		t.Errorf("Power(7, 1) = %d, want 7", got)
	}
	if got := Power(5, 0); got != 1 {
		t.Errorf("Power(5, 0) = %d, want 1", got)
	}
}
