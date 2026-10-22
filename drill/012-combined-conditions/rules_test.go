package rules

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
	if got := FizzBuzz(1); got != "1" {
		t.Errorf("FizzBuzz(1) = %q, want %q", got, "1")
	}
}

func TestIsPrime(t *testing.T) {
	if got := IsPrime(2); got != true {
		t.Errorf("IsPrime(2) = %v, want true", got)
	}
	if got := IsPrime(3); got != true {
		t.Errorf("IsPrime(3) = %v, want true", got)
	}
	if got := IsPrime(17); got != true {
		t.Errorf("IsPrime(17) = %v, want true", got)
	}
	if got := IsPrime(97); got != true {
		t.Errorf("IsPrime(97) = %v, want true", got)
	}
	if got := IsPrime(4); got != false {
		t.Errorf("IsPrime(4) = %v, want false", got)
	}
	if got := IsPrime(9); got != false {
		t.Errorf("IsPrime(9) = %v, want false", got)
	}
	if got := IsPrime(21); got != false {
		t.Errorf("IsPrime(21) = %v, want false", got)
	}
	if got := IsPrime(1); got != false {
		t.Errorf("IsPrime(1) = %v, want false", got)
	}
	if got := IsPrime(0); got != false {
		t.Errorf("IsPrime(0) = %v, want false", got)
	}
}

func TestNextPrime(t *testing.T) {
	if got := NextPrime(10); got != 11 {
		t.Errorf("NextPrime(10) = %d, want 11", got)
	}
	if got := NextPrime(13); got != 17 {
		t.Errorf("NextPrime(13) = %d, want 17", got)
	}
	if got := NextPrime(1); got != 2 {
		t.Errorf("NextPrime(1) = %d, want 2", got)
	}
	if got := NextPrime(0); got != 2 {
		t.Errorf("NextPrime(0) = %d, want 2", got)
	}
	if got := NextPrime(20); got != 23 {
		t.Errorf("NextPrime(20) = %d, want 23", got)
	}
}
