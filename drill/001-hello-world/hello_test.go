package hello

// You don't need to understand everything in this file yet.
// Writing tests is a topic later in the roadmap.
// For now: each TestXxx function calls your code and compares
// what it got with what it wanted.

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World!"

	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestGreet(t *testing.T) {
	got := Greet("Gopher")
	want := "Hello, Gopher!"

	if got != want {
		t.Errorf("Greet(%q) = %q, want %q", "Gopher", got, want)
	}

	got = Greet("Ada")
	want = "Hello, Ada!"

	if got != want {
		t.Errorf("Greet(%q) = %q, want %q", "Ada", got, want)
	}
}

func TestShout(t *testing.T) {
	got := Shout("Gopher")
	want := "HELLO, GOPHER!"

	if got != want {
		t.Errorf("Shout(%q) = %q, want %q", "Gopher", got, want)
	}

	got = Shout("go")
	want = "HELLO, GO!"

	if got != want {
		t.Errorf("Shout(%q) = %q, want %q", "go", got, want)
	}
}
