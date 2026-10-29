package redo

import "testing"

func TestDigitalRoot(t *testing.T) {
	if got := DigitalRoot(9875); got != 2 {
		t.Errorf("DigitalRoot(9875) = %d, want 2", got)
	}
	if got := DigitalRoot(99); got != 9 {
		t.Errorf("DigitalRoot(99) = %d, want 9", got)
	}
	if got := DigitalRoot(10); got != 1 {
		t.Errorf("DigitalRoot(10) = %d, want 1", got)
	}
	if got := DigitalRoot(7); got != 7 {
		t.Errorf("DigitalRoot(7) = %d, want 7", got)
	}
	if got := DigitalRoot(0); got != 0 {
		t.Errorf("DigitalRoot(0) = %d, want 0", got)
	}
}
