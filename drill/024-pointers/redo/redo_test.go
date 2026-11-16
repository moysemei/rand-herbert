package redo

import "testing"

func TestIsAnagram(t *testing.T) {
	if got := IsAnagram("listen", "silent"); got != true {
		t.Errorf("IsAnagram(%q, %q) = %v, want true", "listen", "silent", got)
	}
	if got := IsAnagram("aab", "abb"); got != false {
		t.Errorf("IsAnagram(%q, %q) = %v, want false (same letters, different counts)", "aab", "abb", got)
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
