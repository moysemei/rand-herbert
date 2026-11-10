package redo

import (
	"maps"
	"testing"
)

func TestWordCount(t *testing.T) {
	if got := WordCount("go is fun go"); !maps.Equal(got, map[string]int{"go": 2, "is": 1, "fun": 1}) {
		t.Errorf("WordCount(%q) = %v, want map[fun:1 go:2 is:1]", "go is fun go", got)
	}
	if got := WordCount("Go GO go"); !maps.Equal(got, map[string]int{"go": 3}) {
		t.Errorf("WordCount(%q) = %v, want map[go:3]", "Go GO go", got)
	}
	if got := WordCount("  one  "); !maps.Equal(got, map[string]int{"one": 1}) {
		t.Errorf("WordCount(%q) = %v, want map[one:1]", "  one  ", got)
	}
	if got := WordCount(""); len(got) != 0 {
		t.Errorf("WordCount(%q) = %v, want map[]", "", got)
	}
}
