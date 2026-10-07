package convert

import "testing"

func TestCelsiusToFahrenheit(t *testing.T) {
	if got := CelsiusToFahrenheit(0); got != 32 {
		t.Errorf("CelsiusToFahrenheit(0) = %v, want 32", got)
	}
	if got := CelsiusToFahrenheit(100); got != 212 {
		t.Errorf("CelsiusToFahrenheit(100) = %v, want 212", got)
	}
	if got := CelsiusToFahrenheit(25); got != 77 {
		t.Errorf("CelsiusToFahrenheit(25) = %v, want 77", got)
	}
	if got := CelsiusToFahrenheit(-40); got != -40 {
		t.Errorf("CelsiusToFahrenheit(-40) = %v, want -40", got)
	}
}

func TestMinutesToHours(t *testing.T) {
	if got := MinutesToHours(120); got != 2 {
		t.Errorf("MinutesToHours(120) = %v, want 2", got)
	}
	if got := MinutesToHours(90); got != 1.5 {
		t.Errorf("MinutesToHours(90) = %v, want 1.5", got)
	}
	if got := MinutesToHours(45); got != 0.75 {
		t.Errorf("MinutesToHours(45) = %v, want 0.75", got)
	}
	if got := MinutesToHours(0); got != 0 {
		t.Errorf("MinutesToHours(0) = %v, want 0", got)
	}
}

func TestDescribe(t *testing.T) {
	if got := Describe("Ada", 36); got != "Ada is 36 years old" {
		t.Errorf("Describe(%q, 36) = %q, want %q", "Ada", got, "Ada is 36 years old")
	}
	if got := Describe("Rob", 7); got != "Rob is 7 years old" {
		t.Errorf("Describe(%q, 7) = %q, want %q", "Rob", got, "Rob is 7 years old")
	}
}
