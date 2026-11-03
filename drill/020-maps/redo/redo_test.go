package redo

import (
	"slices"
	"testing"
)

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
	ReverseInPlace(empty)
	if len(empty) != 0 {
		t.Errorf("after ReverseInPlace, [] = %v, want []", empty)
	}
}
