package text

import "testing"

func TestNormalize(t *testing.T) {
	if got := Normalize("A Man A Plan"); got != "amanaplan" {
		t.Errorf("Normalize(%q) = %q, want %q", "A Man A Plan", got, "amanaplan")
	}
	if got := Normalize("  Go  "); got != "go" {
		t.Errorf("Normalize(%q) = %q, want %q", "  Go  ", got, "go")
	}
	if got := Normalize("GOPHER"); got != "gopher" {
		t.Errorf("Normalize(%q) = %q, want %q", "GOPHER", got, "gopher")
	}
	if got := Normalize(""); got != "" {
		t.Errorf("Normalize(%q) = %q, want %q", "", got, "")
	}
}

func TestIsPalindrome(t *testing.T) {
	if got := IsPalindrome("Ana"); got != true {
		t.Errorf("IsPalindrome(%q) = %v, want true", "Ana", got)
	}
	if got := IsPalindrome("A man a plan a canal Panama"); got != true {
		t.Errorf("IsPalindrome(%q) = %v, want true", "A man a plan a canal Panama", got)
	}
	if got := IsPalindrome("socorram me subi no onibus em marrocos"); got != true {
		t.Errorf("IsPalindrome(%q) = %v, want true", "socorram me subi no onibus em marrocos", got)
	}
	if got := IsPalindrome("Go"); got != false {
		t.Errorf("IsPalindrome(%q) = %v, want false", "Go", got)
	}
	if got := IsPalindrome("ab ca"); got != false {
		t.Errorf("IsPalindrome(%q) = %v, want false", "ab ca", got)
	}
	if got := IsPalindrome(""); got != true {
		t.Errorf("IsPalindrome(%q) = %v, want true", "", got)
	}
}

func TestSumNumbers(t *testing.T) {
	if sum, ok := SumNumbers("1 2 3"); sum != 6 || ok != true {
		t.Errorf("SumNumbers(%q) = %d, %v, want 6, true", "1 2 3", sum, ok)
	}
	if sum, ok := SumNumbers(" 10   20 "); sum != 30 || ok != true {
		t.Errorf("SumNumbers(%q) = %d, %v, want 30, true", " 10   20 ", sum, ok)
	}
	if sum, ok := SumNumbers("-2 5"); sum != 3 || ok != true {
		t.Errorf("SumNumbers(%q) = %d, %v, want 3, true", "-2 5", sum, ok)
	}
	if sum, ok := SumNumbers(""); sum != 0 || ok != true {
		t.Errorf("SumNumbers(%q) = %d, %v, want 0, true", "", sum, ok)
	}
	if sum, ok := SumNumbers("4 x 5"); sum != 0 || ok != false {
		t.Errorf("SumNumbers(%q) = %d, %v, want 0, false", "4 x 5", sum, ok)
	}
	if sum, ok := SumNumbers("7 8 nine"); sum != 0 || ok != false {
		t.Errorf("SumNumbers(%q) = %d, %v, want 0, false", "7 8 nine", sum, ok)
	}
}
