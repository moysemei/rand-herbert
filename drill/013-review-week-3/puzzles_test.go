package puzzles

import "testing"

func TestCountPrimes(t *testing.T) {
	if got := CountPrimes(10); got != 4 {
		t.Errorf("CountPrimes(10) = %d, want 4", got)
	}
	if got := CountPrimes(2); got != 1 {
		t.Errorf("CountPrimes(2) = %d, want 1", got)
	}
	if got := CountPrimes(1); got != 0 {
		t.Errorf("CountPrimes(1) = %d, want 0", got)
	}
	if got := CountPrimes(0); got != 0 {
		t.Errorf("CountPrimes(0) = %d, want 0", got)
	}
	if got := CountPrimes(30); got != 10 {
		t.Errorf("CountPrimes(30) = %d, want 10", got)
	}
	if got := CountPrimes(100); got != 25 {
		t.Errorf("CountPrimes(100) = %d, want 25", got)
	}
}

func TestIsPalindromeNumber(t *testing.T) {
	if got := IsPalindromeNumber(121); got != true {
		t.Errorf("IsPalindromeNumber(121) = %v, want true", got)
	}
	if got := IsPalindromeNumber(1221); got != true {
		t.Errorf("IsPalindromeNumber(1221) = %v, want true", got)
	}
	if got := IsPalindromeNumber(7); got != true {
		t.Errorf("IsPalindromeNumber(7) = %v, want true", got)
	}
	if got := IsPalindromeNumber(0); got != true {
		t.Errorf("IsPalindromeNumber(0) = %v, want true", got)
	}
	if got := IsPalindromeNumber(123); got != false {
		t.Errorf("IsPalindromeNumber(123) = %v, want false", got)
	}
	if got := IsPalindromeNumber(10); got != false {
		t.Errorf("IsPalindromeNumber(10) = %v, want false", got)
	}
	if got := IsPalindromeNumber(1210); got != false {
		t.Errorf("IsPalindromeNumber(1210) = %v, want false", got)
	}
}

func TestDigitalRoot(t *testing.T) {
	if got := DigitalRoot(9875); got != 2 {
		t.Errorf("DigitalRoot(9875) = %d, want 2", got)
	}
	if got := DigitalRoot(99); got != 9 {
		t.Errorf("DigitalRoot(99) = %d, want 9", got)
	}
	if got := DigitalRoot(10); got != 1 {
		t.Errorf("DigitalRoot(10) = %d, want 1", got)
	}
	if got := DigitalRoot(7); got != 7 {
		t.Errorf("DigitalRoot(7) = %d, want 7", got)
	}
	if got := DigitalRoot(0); got != 0 {
		t.Errorf("DigitalRoot(0) = %d, want 0", got)
	}
}
