package runes

import "testing"

func TestCountVowels(t *testing.T) {
	if got := CountVowels("banana"); got != 3 {
		t.Errorf("CountVowels(%q) = %d, want 3", "banana", got)
	}
	if got := CountVowels("hello"); got != 2 {
		t.Errorf("CountVowels(%q) = %d, want 2", "hello", got)
	}
	if got := CountVowels("rhythm"); got != 0 {
		t.Errorf("CountVowels(%q) = %d, want 0", "rhythm", got)
	}
	if got := CountVowels(""); got != 0 {
		t.Errorf("CountVowels(%q) = %d, want 0", "", got)
	}
	if got := CountVowels("education"); got != 5 {
		t.Errorf("CountVowels(%q) = %d, want 5", "education", got)
	}
}

func TestSizes(t *testing.T) {
	if b, r := Sizes("go"); b != 2 || r != 2 {
		t.Errorf("Sizes(%q) = %d, %d, want 2, 2", "go", b, r)
	}
	if b, r := Sizes("ação"); b != 6 || r != 4 {
		t.Errorf("Sizes(%q) = %d, %d, want 6, 4", "ação", b, r)
	}
	if b, r := Sizes("日本"); b != 6 || r != 2 {
		t.Errorf("Sizes(%q) = %d, %d, want 6, 2", "日本", b, r)
	}
	if b, r := Sizes(""); b != 0 || r != 0 {
		t.Errorf("Sizes(%q) = %d, %d, want 0, 0", "", b, r)
	}
}

func TestReverse(t *testing.T) {
	if got := Reverse("hello"); got != "olleh" {
		t.Errorf("Reverse(%q) = %q, want %q", "hello", got, "olleh")
	}
	if got := Reverse("ação"); got != "oãça" {
		t.Errorf("Reverse(%q) = %q, want %q", "ação", got, "oãça")
	}
	if got := Reverse("日本語"); got != "語本日" {
		t.Errorf("Reverse(%q) = %q, want %q", "日本語", got, "語本日")
	}
	if got := Reverse("a"); got != "a" {
		t.Errorf("Reverse(%q) = %q, want %q", "a", got, "a")
	}
	if got := Reverse(""); got != "" {
		t.Errorf("Reverse(%q) = %q, want %q", "", got, "")
	}
}
