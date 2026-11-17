//go:build extra

// The extra function of drill 025. Run with: go test -v -tags extra
package pointers2

import "testing"

func TestIsBalanced(t *testing.T) {
	if got := IsBalanced("([]{})"); got != true {
		t.Errorf("IsBalanced(%q) = %v, want true", "([]{})", got)
	}
	if got := IsBalanced("a(b)c"); got != true {
		t.Errorf("IsBalanced(%q) = %v, want true (other characters are ignored)", "a(b)c", got)
	}
	if got := IsBalanced("([)]"); got != false {
		t.Errorf("IsBalanced(%q) = %v, want false (the ) closes before the ])", "([)]", got)
	}
	if got := IsBalanced("(("); got != false {
		t.Errorf("IsBalanced(%q) = %v, want false (left open)", "((", got)
	}
	if got := IsBalanced("())"); got != false {
		t.Errorf("IsBalanced(%q) = %v, want false (closed without opening)", "())", got)
	}
	if got := IsBalanced("}"); got != false {
		t.Errorf("IsBalanced(%q) = %v, want false", "}", got)
	}
	if got := IsBalanced(""); got != true {
		t.Errorf("IsBalanced(%q) = %v, want true", "", got)
	}
}
