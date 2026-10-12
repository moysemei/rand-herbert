package numbers

import "testing"

func TestIsEven(t *testing.T) {
	if got := IsEven(4); got != true {
		t.Errorf("IsEven(4) = %v, want true", got)
	}
	if got := IsEven(7); got != false {
		t.Errorf("IsEven(7) = %v, want false", got)
	}
	if got := IsEven(0); got != true {
		t.Errorf("IsEven(0) = %v, want true", got)
	}
	if got := IsEven(-3); got != false {
		t.Errorf("IsEven(-3) = %v, want false", got)
	}
}

func TestSecondsToClock(t *testing.T) {
	if got := SecondsToClock(59); got != "0h 0m 59s" {
		t.Errorf("SecondsToClock(59) = %q, want %q", got, "0h 0m 59s")
	}
	if got := SecondsToClock(3665); got != "1h 1m 5s" {
		t.Errorf("SecondsToClock(3665) = %q, want %q", got, "1h 1m 5s")
	}
	if got := SecondsToClock(7200); got != "2h 0m 0s" {
		t.Errorf("SecondsToClock(7200) = %q, want %q", got, "2h 0m 0s")
	}
	if got := SecondsToClock(86399); got != "23h 59m 59s" {
		t.Errorf("SecondsToClock(86399) = %q, want %q", got, "23h 59m 59s")
	}
}

func TestIsLeapYear(t *testing.T) {
	if got := IsLeapYear(2024); got != true {
		t.Errorf("IsLeapYear(2024) = %v, want true", got)
	}
	if got := IsLeapYear(2023); got != false {
		t.Errorf("IsLeapYear(2023) = %v, want false", got)
	}
	if got := IsLeapYear(1900); got != false {
		t.Errorf("IsLeapYear(1900) = %v, want false", got)
	}
	if got := IsLeapYear(2000); got != true {
		t.Errorf("IsLeapYear(2000) = %v, want true", got)
	}
}
