package redo

import (
	"slices"
	"testing"
)

func TestUnique(t *testing.T) {
	for i := 0; i < 20; i++ { // repeated: a wrong order from map iteration would show up
		if got := Unique([]int{3, 1, 3, 2, 1}); !slices.Equal(got, []int{3, 1, 2}) {
			t.Fatalf("Unique([3 1 3 2 1]) = %v, want [3 1 2]", got)
		}
	}
	if got := Unique([]int{5, 5}); !slices.Equal(got, []int{5}) {
		t.Errorf("Unique([5 5]) = %v, want [5]", got)
	}
	if got := Unique([]int{9, 8, 7}); !slices.Equal(got, []int{9, 8, 7}) {
		t.Errorf("Unique([9 8 7]) = %v, want [9 8 7]", got)
	}
	if got := Unique([]int{}); len(got) != 0 {
		t.Errorf("Unique([]) = %v, want []", got)
	}
}
