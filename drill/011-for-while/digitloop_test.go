package digitloop

import "testing"

func TestSumDigits(t *testing.T) {
	if got := SumDigits(472); got != 13 {
		t.Errorf("SumDigits(472) = %d, want 13", got)
	}
	if got := SumDigits(1000); got != 1 {
		t.Errorf("SumDigits(1000) = %d, want 1", got)
	}
	if got := SumDigits(9); got != 9 {
		t.Errorf("SumDigits(9) = %d, want 9", got)
	}
	if got := SumDigits(0); got != 0 {
		t.Errorf("SumDigits(0) = %d, want 0", got)
	}
	if got := SumDigits(99999); got != 45 {
		t.Errorf("SumDigits(99999) = %d, want 45", got)
	}
}

func TestCountDigits(t *testing.T) {
	if got := CountDigits(472); got != 3 {
		t.Errorf("CountDigits(472) = %d, want 3", got)
	}
	if got := CountDigits(1000); got != 4 {
		t.Errorf("CountDigits(1000) = %d, want 4", got)
	}
	if got := CountDigits(7); got != 1 {
		t.Errorf("CountDigits(7) = %d, want 1", got)
	}
	if got := CountDigits(0); got != 1 {
		t.Errorf("CountDigits(0) = %d, want 1", got)
	}
	if got := CountDigits(123456789); got != 9 {
		t.Errorf("CountDigits(123456789) = %d, want 9", got)
	}
}

func TestReverseNumber(t *testing.T) {
	if got := ReverseNumber(472); got != 274 {
		t.Errorf("ReverseNumber(472) = %d, want 274", got)
	}
	if got := ReverseNumber(1200); got != 21 {
		t.Errorf("ReverseNumber(1200) = %d, want 21", got)
	}
	if got := ReverseNumber(123456); got != 654321 {
		t.Errorf("ReverseNumber(123456) = %d, want 654321", got)
	}
	if got := ReverseNumber(5); got != 5 {
		t.Errorf("ReverseNumber(5) = %d, want 5", got)
	}
	if got := ReverseNumber(0); got != 0 {
		t.Errorf("ReverseNumber(0) = %d, want 0", got)
	}
}
