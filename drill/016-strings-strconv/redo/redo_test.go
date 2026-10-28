package redo

import "testing"

func TestFizzBuzz(t *testing.T) {
	if got := FizzBuzz(15); got != "FizzBuzz" {
		t.Errorf("FizzBuzz(15) = %q, want %q", got, "FizzBuzz")
	}
	if got := FizzBuzz(30); got != "FizzBuzz" {
		t.Errorf("FizzBuzz(30) = %q, want %q", got, "FizzBuzz")
	}
	if got := FizzBuzz(9); got != "Fizz" {
		t.Errorf("FizzBuzz(9) = %q, want %q", got, "Fizz")
	}
	if got := FizzBuzz(10); got != "Buzz" {
		t.Errorf("FizzBuzz(10) = %q, want %q", got, "Buzz")
	}
	if got := FizzBuzz(7); got != "7" {
		t.Errorf("FizzBuzz(7) = %q, want %q", got, "7")
	}
}
