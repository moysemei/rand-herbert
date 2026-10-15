package calendar

import "testing"

func TestDayType(t *testing.T) {
	if got := DayType("saturday"); got != "weekend" {
		t.Errorf("DayType(%q) = %q, want %q", "saturday", got, "weekend")
	}
	if got := DayType("sunday"); got != "weekend" {
		t.Errorf("DayType(%q) = %q, want %q", "sunday", got, "weekend")
	}
	if got := DayType("monday"); got != "weekday" {
		t.Errorf("DayType(%q) = %q, want %q", "monday", got, "weekday")
	}
	if got := DayType("wednesday"); got != "weekday" {
		t.Errorf("DayType(%q) = %q, want %q", "wednesday", got, "weekday")
	}
	if got := DayType("friday"); got != "weekday" {
		t.Errorf("DayType(%q) = %q, want %q", "friday", got, "weekday")
	}
	if got := DayType("funday"); got != "invalid" {
		t.Errorf("DayType(%q) = %q, want %q", "funday", got, "invalid")
	}
}

func TestGrade(t *testing.T) {
	if got := Grade(95); got != "A" {
		t.Errorf("Grade(95) = %q, want %q", got, "A")
	}
	if got := Grade(100); got != "A" {
		t.Errorf("Grade(100) = %q, want %q", got, "A")
	}
	if got := Grade(90); got != "A" {
		t.Errorf("Grade(90) = %q, want %q", got, "A")
	}
	if got := Grade(89); got != "B" {
		t.Errorf("Grade(89) = %q, want %q", got, "B")
	}
	if got := Grade(75); got != "C" {
		t.Errorf("Grade(75) = %q, want %q", got, "C")
	}
	if got := Grade(60); got != "D" {
		t.Errorf("Grade(60) = %q, want %q", got, "D")
	}
	if got := Grade(59); got != "F" {
		t.Errorf("Grade(59) = %q, want %q", got, "F")
	}
	if got := Grade(0); got != "F" {
		t.Errorf("Grade(0) = %q, want %q", got, "F")
	}
	if got := Grade(101); got != "invalid" {
		t.Errorf("Grade(101) = %q, want %q", got, "invalid")
	}
	if got := Grade(-1); got != "invalid" {
		t.Errorf("Grade(-1) = %q, want %q", got, "invalid")
	}
}

func TestDaysInMonth(t *testing.T) {
	if got := DaysInMonth(1, 2023); got != 31 {
		t.Errorf("DaysInMonth(1, 2023) = %d, want 31", got)
	}
	if got := DaysInMonth(4, 2023); got != 30 {
		t.Errorf("DaysInMonth(4, 2023) = %d, want 30", got)
	}
	if got := DaysInMonth(11, 2023); got != 30 {
		t.Errorf("DaysInMonth(11, 2023) = %d, want 30", got)
	}
	if got := DaysInMonth(12, 2023); got != 31 {
		t.Errorf("DaysInMonth(12, 2023) = %d, want 31", got)
	}
	if got := DaysInMonth(2, 2023); got != 28 {
		t.Errorf("DaysInMonth(2, 2023) = %d, want 28", got)
	}
	if got := DaysInMonth(2, 2024); got != 29 {
		t.Errorf("DaysInMonth(2, 2024) = %d, want 29", got)
	}
	if got := DaysInMonth(2, 1900); got != 28 {
		t.Errorf("DaysInMonth(2, 1900) = %d, want 28", got)
	}
	if got := DaysInMonth(2, 2000); got != 29 {
		t.Errorf("DaysInMonth(2, 2000) = %d, want 29", got)
	}
	if got := DaysInMonth(13, 2023); got != 0 {
		t.Errorf("DaysInMonth(13, 2023) = %d, want 0", got)
	}
	if got := DaysInMonth(0, 2023); got != 0 {
		t.Errorf("DaysInMonth(0, 2023) = %d, want 0", got)
	}
}
