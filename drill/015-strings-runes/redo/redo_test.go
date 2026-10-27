package redo

import "testing"

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
