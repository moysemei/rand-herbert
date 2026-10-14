package decide

import "testing"

func TestMax(t *testing.T) {
	if got := Max(3, 7); got != 7 {
		t.Errorf("Max(3, 7) = %d, want 7", got)
	}
	if got := Max(7, 3); got != 7 {
		t.Errorf("Max(7, 3) = %d, want 7", got)
	}
	if got := Max(-2, -5); got != -2 {
		t.Errorf("Max(-2, -5) = %d, want -2", got)
	}
	if got := Max(4, 4); got != 4 {
		t.Errorf("Max(4, 4) = %d, want 4", got)
	}
}

func TestClassify(t *testing.T) {
	if got := Classify(-4); got != "negative" {
		t.Errorf("Classify(-4) = %q, want %q", got, "negative")
	}
	if got := Classify(-3); got != "negative" {
		t.Errorf("Classify(-3) = %q, want %q", got, "negative")
	}
	if got := Classify(0); got != "zero" {
		t.Errorf("Classify(0) = %q, want %q", got, "zero")
	}
	if got := Classify(8); got != "positive even" {
		t.Errorf("Classify(8) = %q, want %q", got, "positive even")
	}
	if got := Classify(7); got != "positive odd" {
		t.Errorf("Classify(7) = %q, want %q", got, "positive odd")
	}
}

func TestRoundToTen(t *testing.T) {
	if got := RoundToTen(47); got != 50 {
		t.Errorf("RoundToTen(47) = %d, want 50", got)
	}
	if got := RoundToTen(43); got != 40 {
		t.Errorf("RoundToTen(43) = %d, want 40", got)
	}
	if got := RoundToTen(45); got != 50 {
		t.Errorf("RoundToTen(45) = %d, want 50", got)
	}
	if got := RoundToTen(50); got != 50 {
		t.Errorf("RoundToTen(50) = %d, want 50", got)
	}
	if got := RoundToTen(0); got != 0 {
		t.Errorf("RoundToTen(0) = %d, want 0", got)
	}
	if got := RoundToTen(994); got != 990 {
		t.Errorf("RoundToTen(994) = %d, want 990", got)
	}
}
