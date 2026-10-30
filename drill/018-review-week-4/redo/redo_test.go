package redo

import "testing"

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
	if got := Reverse(""); got != "" {
		t.Errorf("Reverse(%q) = %q, want %q", "", got, "")
	}
}
