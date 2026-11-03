package counting

import (
	"maps"
	"slices"
	"testing"
)

func TestAddStock(t *testing.T) {
	stock := map[string]int{"apple": 3}
	AddStock(stock, "apple", 2)
	if !maps.Equal(stock, map[string]int{"apple": 5}) {
		t.Errorf("after AddStock(apple, 2), stock = %v, want map[apple:5]", stock)
	}
	AddStock(stock, "pear", 4)
	if !maps.Equal(stock, map[string]int{"apple": 5, "pear": 4}) {
		t.Errorf("after AddStock(pear, 4), stock = %v, want map[apple:5 pear:4]", stock)
	}
}

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

func TestUnique(t *testing.T) {
	for i := 0; i < 20; i++ { // repeated: a wrong order from map iteration would show up
		if got := Unique([]int{3, 1, 3, 2, 1}); !slices.Equal(got, []int{3, 1, 2}) {
			t.Fatalf("Unique([3 1 3 2 1]) = %v, want [3 1 2]", got)
		}
	}
	if got := Unique([]int{5, 5}); !slices.Equal(got, []int{5}) {
		t.Errorf("Unique([5 5]) = %v, want [5]", got)
	}
	if got := Unique([]int{9, 8, 7}); !slices.Equal(got, []int{9, 8, 7}) {
		t.Errorf("Unique([9 8 7]) = %v, want [9 8 7]", got)
	}
	if got := Unique([]int{}); len(got) != 0 {
		t.Errorf("Unique([]) = %v, want []", got)
	}
}

func TestFirstRepeated(t *testing.T) {
	if r, ok := FirstRepeated("abcba"); r != 'b' || ok != true {
		t.Errorf("FirstRepeated(%q) = %q, %v, want 'b', true", "abcba", r, ok)
	}
	if r, ok := FirstRepeated("hello"); r != 'l' || ok != true {
		t.Errorf("FirstRepeated(%q) = %q, %v, want 'l', true", "hello", r, ok)
	}
	if r, ok := FirstRepeated("ação ã"); r != 'ã' || ok != true {
		t.Errorf("FirstRepeated(%q) = %q, %v, want 'ã', true", "ação ã", r, ok)
	}
	if r, ok := FirstRepeated("abc"); r != 0 || ok != false {
		t.Errorf("FirstRepeated(%q) = %q, %v, want 0, false", "abc", r, ok)
	}
	if r, ok := FirstRepeated(""); r != 0 || ok != false {
		t.Errorf("FirstRepeated(%q) = %q, %v, want 0, false", "", r, ok)
	}
}
