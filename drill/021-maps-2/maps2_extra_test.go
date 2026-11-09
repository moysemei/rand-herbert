//go:build extra

// The two extra functions of drill 021. Run with: go test -v -tags extra

package maps2

import (
	"maps"
	"testing"
)

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
