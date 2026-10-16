package parking

import "testing"

func TestFormatDuration(t *testing.T) {
	if got := FormatDuration(125); got != "2h 5m" {
		t.Errorf("FormatDuration(125) = %q, want %q", got, "2h 5m")
	}
	if got := FormatDuration(59); got != "0h 59m" {
		t.Errorf("FormatDuration(59) = %q, want %q", got, "0h 59m")
	}
	if got := FormatDuration(60); got != "1h 0m" {
		t.Errorf("FormatDuration(60) = %q, want %q", got, "1h 0m")
	}
	if got := FormatDuration(0); got != "0h 0m" {
		t.Errorf("FormatDuration(0) = %q, want %q", got, "0h 0m")
	}
}

func TestBilledHours(t *testing.T) {
	if got := BilledHours(0); got != 0 {
		t.Errorf("BilledHours(0) = %d, want 0", got)
	}
	if got := BilledHours(1); got != 1 {
		t.Errorf("BilledHours(1) = %d, want 1", got)
	}
	if got := BilledHours(60); got != 1 {
		t.Errorf("BilledHours(60) = %d, want 1", got)
	}
	if got := BilledHours(61); got != 2 {
		t.Errorf("BilledHours(61) = %d, want 2", got)
	}
	if got := BilledHours(125); got != 3 {
		t.Errorf("BilledHours(125) = %d, want 3", got)
	}
	if got := BilledHours(180); got != 3 {
		t.Errorf("BilledHours(180) = %d, want 3", got)
	}
}

func TestPrice(t *testing.T) {
	if got := Price(10); got != 0 {
		t.Errorf("Price(10) = %d, want 0", got)
	}
	if got := Price(15); got != 0 {
		t.Errorf("Price(15) = %d, want 0", got)
	}
	if got := Price(16); got != 8 {
		t.Errorf("Price(16) = %d, want 8", got)
	}
	if got := Price(60); got != 8 {
		t.Errorf("Price(60) = %d, want 8", got)
	}
	if got := Price(61); got != 13 {
		t.Errorf("Price(61) = %d, want 13", got)
	}
	if got := Price(125); got != 18 {
		t.Errorf("Price(125) = %d, want 18", got)
	}
	if got := Price(181); got != 23 {
		t.Errorf("Price(181) = %d, want 23", got)
	}
}
