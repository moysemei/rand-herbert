package redo

import "testing"

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
