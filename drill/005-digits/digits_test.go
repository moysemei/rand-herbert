package digits

import "testing"

func TestLastDigit(t *testing.T) {
	if got := LastDigit(1234); got != 4 {
		t.Errorf("LastDigit(1234) = %d, want 4", got)
	}
	if got := LastDigit(7); got != 7 {
		t.Errorf("LastDigit(7) = %d, want 7", got)
	}
	if got := LastDigit(90); got != 0 {
		t.Errorf("LastDigit(90) = %d, want 0", got)
	}
}

func TestDescribeDigits(t *testing.T) {
	if got := DescribeDigits(472); got != "hundreds: 4, tens: 7, ones: 2" {
		t.Errorf("DescribeDigits(472) = %q, want %q", got, "hundreds: 4, tens: 7, ones: 2")
	}
	if got := DescribeDigits(905); got != "hundreds: 9, tens: 0, ones: 5" {
		t.Errorf("DescribeDigits(905) = %q, want %q", got, "hundreds: 9, tens: 0, ones: 5")
	}
	if got := DescribeDigits(100); got != "hundreds: 1, tens: 0, ones: 0" {
		t.Errorf("DescribeDigits(100) = %q, want %q", got, "hundreds: 1, tens: 0, ones: 0")
	}
	if got := DescribeDigits(999); got != "hundreds: 9, tens: 9, ones: 9" {
		t.Errorf("DescribeDigits(999) = %q, want %q", got, "hundreds: 9, tens: 9, ones: 9")
	}
}

func TestReverseThreeDigits(t *testing.T) {
	if got := ReverseThreeDigits(472); got != 274 {
		t.Errorf("ReverseThreeDigits(472) = %d, want 274", got)
	}
	if got := ReverseThreeDigits(905); got != 509 {
		t.Errorf("ReverseThreeDigits(905) = %d, want 509", got)
	}
	if got := ReverseThreeDigits(120); got != 21 {
		t.Errorf("ReverseThreeDigits(120) = %d, want 21", got)
	}
	if got := ReverseThreeDigits(111); got != 111 {
		t.Errorf("ReverseThreeDigits(111) = %d, want 111", got)
	}
}
