package redo

import "testing"

func TestSumNumbers(t *testing.T) {
	if sum, ok := SumNumbers("1 2 3"); sum != 6 || ok != true {
		t.Errorf("SumNumbers(%q) = %d, %v, want 6, true", "1 2 3", sum, ok)
	}
	if sum, ok := SumNumbers(" 10   20 "); sum != 30 || ok != true {
		t.Errorf("SumNumbers(%q) = %d, %v, want 30, true", " 10   20 ", sum, ok)
	}
	if sum, ok := SumNumbers("-2 5"); sum != 3 || ok != true {
		t.Errorf("SumNumbers(%q) = %d, %v, want 3, true", "-2 5", sum, ok)
	}
	if sum, ok := SumNumbers(""); sum != 0 || ok != true {
		t.Errorf("SumNumbers(%q) = %d, %v, want 0, true", "", sum, ok)
	}
	if sum, ok := SumNumbers("4 x 5"); sum != 0 || ok != false {
		t.Errorf("SumNumbers(%q) = %d, %v, want 0, false", "4 x 5", sum, ok)
	}
}
