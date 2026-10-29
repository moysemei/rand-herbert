package lists

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3}); got != 6 {
		t.Errorf("Sum([1 2 3]) = %d, want 6", got)
	}
	if got := Sum([]int{-5, 5}); got != 0 {
		t.Errorf("Sum([-5 5]) = %d, want 0", got)
	}
	if got := Sum([]int{10}); got != 10 {
		t.Errorf("Sum([10]) = %d, want 10", got)
	}
	if got := Sum([]int{}); got != 0 {
		t.Errorf("Sum([]) = %d, want 0", got)
	}
}

func TestMaxOf(t *testing.T) {
	if got, ok := MaxOf([]int{3, 9, 2}); got != 9 || ok != true {
		t.Errorf("MaxOf([3 9 2]) = %d, %v, want 9, true", got, ok)
	}
	if got, ok := MaxOf([]int{-3, -1, -7}); got != -1 || ok != true {
		t.Errorf("MaxOf([-3 -1 -7]) = %d, %v, want -1, true", got, ok)
	}
	if got, ok := MaxOf([]int{5}); got != 5 || ok != true {
		t.Errorf("MaxOf([5]) = %d, %v, want 5, true", got, ok)
	}
	if got, ok := MaxOf([]int{2, 8, 8, 1}); got != 8 || ok != true {
		t.Errorf("MaxOf([2 8 8 1]) = %d, %v, want 8, true", got, ok)
	}
	if got, ok := MaxOf([]int{}); got != 0 || ok != false {
		t.Errorf("MaxOf([]) = %d, %v, want 0, false", got, ok)
	}
}

func TestCountdown(t *testing.T) {
	if got := Countdown(3); !slices.Equal(got, []int{3, 2, 1}) {
		t.Errorf("Countdown(3) = %v, want [3 2 1]", got)
	}
	if got := Countdown(1); !slices.Equal(got, []int{1}) {
		t.Errorf("Countdown(1) = %v, want [1]", got)
	}
	if got := Countdown(0); len(got) != 0 {
		t.Errorf("Countdown(0) = %v, want []", got)
	}
}

func TestReverseInPlace(t *testing.T) {
	even := []int{1, 2, 3, 4}
	ReverseInPlace(even)
	if !slices.Equal(even, []int{4, 3, 2, 1}) {
		t.Errorf("after ReverseInPlace, [1 2 3 4] = %v, want [4 3 2 1]", even)
	}

	odd := []int{1, 2, 3}
	ReverseInPlace(odd)
	if !slices.Equal(odd, []int{3, 2, 1}) {
		t.Errorf("after ReverseInPlace, [1 2 3] = %v, want [3 2 1]", odd)
	}

	one := []int{7}
	ReverseInPlace(one)
	if !slices.Equal(one, []int{7}) {
		t.Errorf("after ReverseInPlace, [7] = %v, want [7]", one)
	}

	empty := []int{}
	ReverseInPlace(empty) // must not panic
	if len(empty) != 0 {
		t.Errorf("after ReverseInPlace, [] = %v, want []", empty)
	}
}
