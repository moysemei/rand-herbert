package redo

import "testing"

func TestLongestWord(t *testing.T) {
	if got := LongestWord("the quick brown fox"); got != "quick" {
		t.Errorf("LongestWord(%q) = %q, want %q", "the quick brown fox", got, "quick")
	}
	if got := LongestWord("go is fun"); got != "fun" {
		t.Errorf("LongestWord(%q) = %q, want %q", "go is fun", got, "fun")
	}
	if got := LongestWord("  hello  "); got != "hello" {
		t.Errorf("LongestWord(%q) = %q, want %q", "  hello  ", got, "hello")
	}
	if got := LongestWord(""); got != "" {
		t.Errorf("LongestWord(%q) = %q, want %q", "", got, "")
	}
}
