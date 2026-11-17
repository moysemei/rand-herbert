package pointers2

import "testing"

func TestAbs(t *testing.T) {
	if got := Abs(-7); got != 7 {
		t.Errorf("Abs(-7) = %d, want 7", got)
	}
	if got := Abs(3); got != 3 {
		t.Errorf("Abs(3) = %d, want 3", got)
	}
	if got := Abs(0); got != 0 {
		t.Errorf("Abs(0) = %d, want 0", got)
	}
}

func TestMakeAbs(t *testing.T) {
	x := -7
	MakeAbs(&x)
	if x != 7 {
		t.Errorf("after MakeAbs(&x) with x = -7, x = %d, want 7", x)
	}
	y := 3
	MakeAbs(&y)
	if y != 3 {
		t.Errorf("after MakeAbs(&y) with y = 3, y = %d, want 3", y)
	}
}

func TestPush(t *testing.T) {
	var s Stack
	s.Push('a')
	s.Push('b')
	if len(s.items) != 2 || s.items[1] != 'b' {
		t.Errorf("after Push('a'), Push('b'), items = %q, want ['a' 'b']", s.items)
	}
}

func TestPop(t *testing.T) {
	var s Stack
	s.Push('(')
	s.Push('[')
	if r, ok := s.Pop(); r != '[' || ok != true {
		t.Fatalf("1st Pop() after Push('('), Push('[') = %q, %v, want '[', true", r, ok)
	}
	if r, ok := s.Pop(); r != '(' || ok != true {
		t.Fatalf("2nd Pop() = %q, %v, want '(', true", r, ok)
	}
	if r, ok := s.Pop(); r != 0 || ok != false {
		t.Fatalf("Pop() on an empty stack = %q, %v, want 0, false", r, ok)
	}
	s.Push('x')
	if r, ok := s.Pop(); r != 'x' || ok != true {
		t.Errorf("Pop() after emptying and Push('x') = %q, %v, want 'x', true", r, ok)
	}
}
