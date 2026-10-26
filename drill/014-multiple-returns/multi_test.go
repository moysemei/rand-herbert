package multi

import "testing"

func TestSplitDigits(t *testing.T) {
	if h, te, o := SplitDigits(472); h != 4 || te != 7 || o != 2 {
		t.Errorf("SplitDigits(472) = %d, %d, %d, want 4, 7, 2", h, te, o)
	}
	if h, te, o := SplitDigits(905); h != 9 || te != 0 || o != 5 {
		t.Errorf("SplitDigits(905) = %d, %d, %d, want 9, 0, 5", h, te, o)
	}
	if h, te, o := SplitDigits(100); h != 1 || te != 0 || o != 0 {
		t.Errorf("SplitDigits(100) = %d, %d, %d, want 1, 0, 0", h, te, o)
	}
	if h, te, o := SplitDigits(999); h != 9 || te != 9 || o != 9 {
		t.Errorf("SplitDigits(999) = %d, %d, %d, want 9, 9, 9", h, te, o)
	}
}

func TestDivide(t *testing.T) {
	if q, ok := Divide(7, 2); q != 3 || ok != true {
		t.Errorf("Divide(7, 2) = %d, %v, want 3, true", q, ok)
	}
	if q, ok := Divide(6, 3); q != 2 || ok != true {
		t.Errorf("Divide(6, 3) = %d, %v, want 2, true", q, ok)
	}
	if q, ok := Divide(0, 5); q != 0 || ok != true {
		t.Errorf("Divide(0, 5) = %d, %v, want 0, true", q, ok)
	}
	if q, ok := Divide(5, 0); q != 0 || ok != false {
		t.Errorf("Divide(5, 0) = %d, %v, want 0, false", q, ok)
	}
}

func TestMinMaxDigit(t *testing.T) {
	if lo, hi := MinMaxDigit(4729); lo != 2 || hi != 9 {
		t.Errorf("MinMaxDigit(4729) = %d, %d, want 2, 9", lo, hi)
	}
	if lo, hi := MinMaxDigit(1000); lo != 0 || hi != 1 {
		t.Errorf("MinMaxDigit(1000) = %d, %d, want 0, 1", lo, hi)
	}
	if lo, hi := MinMaxDigit(5555); lo != 5 || hi != 5 {
		t.Errorf("MinMaxDigit(5555) = %d, %d, want 5, 5", lo, hi)
	}
	if lo, hi := MinMaxDigit(7); lo != 7 || hi != 7 {
		t.Errorf("MinMaxDigit(7) = %d, %d, want 7, 7", lo, hi)
	}
	if lo, hi := MinMaxDigit(0); lo != 0 || hi != 0 {
		t.Errorf("MinMaxDigit(0) = %d, %d, want 0, 0", lo, hi)
	}
	if lo, hi := MinMaxDigit(80); lo != 0 || hi != 8 {
		t.Errorf("MinMaxDigit(80) = %d, %d, want 0, 8", lo, hi)
	}
}
