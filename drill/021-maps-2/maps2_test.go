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

func TestRemoveItem(t *testing.T) {
	stock := map[string]int{"apple": 3, "pear": 1}
	if got := RemoveItem(stock, "apple"); got != true {
		t.Errorf("RemoveItem(stock, %q) = %v, want true", "apple", got)
	}
	if !maps.Equal(stock, map[string]int{"pear": 1}) {
		t.Errorf("after RemoveItem(apple), stock = %v, want map[pear:1]", stock)
	}
	if got := RemoveItem(stock, "kiwi"); got != false {
		t.Errorf("RemoveItem(stock, %q) = %v, want false", "kiwi", got)
	}
	if !maps.Equal(stock, map[string]int{"pear": 1}) {
		t.Errorf("after RemoveItem(kiwi), stock = %v, want map[pear:1]", stock)
	}
	zero := map[string]int{"water": 0}
	if got := RemoveItem(zero, "water"); got != true {
		t.Errorf("RemoveItem(map[water:0], %q) = %v, want true (it existed, with value 0)", "water", got)
	}
	if len(zero) != 0 {
		t.Errorf("after RemoveItem(water), map = %v, want map[]", zero)
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

func TestIsAnagram(t *testing.T) {
	if got := IsAnagram("listen", "silent"); got != true {
		t.Errorf("IsAnagram(%q, %q) = %v, want true", "listen", "silent", got)
	}
	if got := IsAnagram("aab", "abb"); got != false {
		t.Errorf("IsAnagram(%q, %q) = %v, want false", "aab", "abb", got)
	}
	if got := IsAnagram("rat", "car"); got != false {
		t.Errorf("IsAnagram(%q, %q) = %v, want false", "rat", "car", got)
	}
	if got := IsAnagram("ab", "abc"); got != false {
		t.Errorf("IsAnagram(%q, %q) = %v, want false", "ab", "abc", got)
	}
	if got := IsAnagram("", ""); got != true {
		t.Errorf("IsAnagram(%q, %q) = %v, want true", "", "", got)
	}
}
