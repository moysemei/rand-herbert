package maps2

import (
	"maps"
	"testing"
)

func TestPrice(t *testing.T) {
	menu := map[string]int{"coffee": 6, "water": 0}
	if p, ok := Price(menu, "coffee"); p != 6 || ok != true {
		t.Errorf("Price(menu, %q) = %d, %v, want 6, true", "coffee", p, ok)
	}
	if p, ok := Price(menu, "water"); p != 0 || ok != true {
		t.Errorf("Price(menu, %q) = %d, %v, want 0, true", "water", p, ok)
	}
	if p, ok := Price(menu, "cake"); p != 0 || ok != false {
		t.Errorf("Price(menu, %q) = %d, %v, want 0, false", "cake", p, ok)
	}
}

func TestCharFrequency(t *testing.T) {
	if got := CharFrequency("banana"); !maps.Equal(got, map[rune]int{'a': 3, 'b': 1, 'n': 2}) {
		t.Errorf("CharFrequency(%q) = %v, want a:3 b:1 n:2", "banana", got)
	}
	if got := CharFrequency("ação"); !maps.Equal(got, map[rune]int{'a': 1, 'ç': 1, 'ã': 1, 'o': 1}) {
		t.Errorf("CharFrequency(%q) = %v, want a:1 ç:1 ã:1 o:1", "ação", got)
	}
	if got := CharFrequency("a a"); !maps.Equal(got, map[rune]int{'a': 2, ' ': 1}) {
		t.Errorf("CharFrequency(%q) = %v, want a:2 and one space", "a a", got)
	}
	if got := CharFrequency(""); len(got) != 0 {
		t.Errorf("CharFrequency(%q) = %v, want map[]", "", got)
	}
}
