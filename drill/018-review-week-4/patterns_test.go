package patterns

import (
	"slices"
	"testing"
)

func TestCountChar(t *testing.T) {
	if got := CountChar("banana", 'a'); got != 3 {
		t.Errorf("CountChar(%q, 'a') = %d, want 3", "banana", got)
	}
	if got := CountChar("ação", 'ç'); got != 1 {
		t.Errorf("CountChar(%q, 'ç') = %d, want 1", "ação", got)
	}
	if got := CountChar("go", 'x'); got != 0 {
		t.Errorf("CountChar(%q, 'x') = %d, want 0", "go", got)
	}
	if got := CountChar("", 'a'); got != 0 {
		t.Errorf("CountChar(%q, 'a') = %d, want 0", "", got)
	}
}

func TestContains(t *testing.T) {
	if got := Contains([]int{4, 8, 15}, 8); got != true {
		t.Errorf("Contains([4 8 15], 8) = %v, want true", got)
	}
	if got := Contains([]int{4, 8, 15}, 15); got != true {
		t.Errorf("Contains([4 8 15], 15) = %v, want true", got)
	}
	if got := Contains([]int{4, 8, 15}, 7); got != false {
		t.Errorf("Contains([4 8 15], 7) = %v, want false", got)
	}
	if got := Contains([]int{}, 1); got != false {
		t.Errorf("Contains([], 1) = %v, want false", got)
	}
}

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
	if got := LongestWord("a bb ccc"); got != "ccc" {
		t.Errorf("LongestWord(%q) = %q, want %q", "a bb ccc", got, "ccc")
	}
	if got := LongestWord(""); got != "" {
		t.Errorf("LongestWord(%q) = %q, want %q", "", got, "")
	}
}

func TestParseAll(t *testing.T) {
	if got, ok := ParseAll([]string{"1", "2", "3"}); !slices.Equal(got, []int{1, 2, 3}) || ok != true {
		t.Errorf(`ParseAll(["1" "2" "3"]) = %v, %v, want [1 2 3], true`, got, ok)
	}
	if got, ok := ParseAll([]string{"10", "-4"}); !slices.Equal(got, []int{10, -4}) || ok != true {
		t.Errorf(`ParseAll(["10" "-4"]) = %v, %v, want [10 -4], true`, got, ok)
	}
	if got, ok := ParseAll([]string{"1", "x", "3"}); len(got) != 0 || ok != false {
		t.Errorf(`ParseAll(["1" "x" "3"]) = %v, %v, want [], false`, got, ok)
	}
	if got, ok := ParseAll([]string{}); len(got) != 0 || ok != true {
		t.Errorf(`ParseAll([]) = %v, %v, want [], true`, got, ok)
	}
}
