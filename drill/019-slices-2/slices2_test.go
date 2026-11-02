package slices2

import (
	"slices"
	"testing"
)

func TestEvens(t *testing.T) {
	input := []int{1, 2, 3, 4}
	if got := Evens(input); !slices.Equal(got, []int{2, 4}) {
		t.Errorf("Evens([1 2 3 4]) = %v, want [2 4]", got)
	}
	if !slices.Equal(input, []int{1, 2, 3, 4}) {
		t.Errorf("Evens changed its input: it is now %v, want [1 2 3 4]", input)
	}
	if got := Evens([]int{1, 3}); len(got) != 0 {
		t.Errorf("Evens([1 3]) = %v, want []", got)
	}
	if got := Evens([]int{}); len(got) != 0 {
		t.Errorf("Evens([]) = %v, want []", got)
	}
	if got := Evens([]int{-2, 0, 5}); !slices.Equal(got, []int{-2, 0}) {
		t.Errorf("Evens([-2 0 5]) = %v, want [-2 0]", got)
	}
}

func TestFirstN(t *testing.T) {
	if got := FirstN([]int{1, 2, 3, 4}, 2); !slices.Equal(got, []int{1, 2}) {
		t.Errorf("FirstN([1 2 3 4], 2) = %v, want [1 2]", got)
	}
	if got := FirstN([]int{1, 2, 3, 4}, 10); !slices.Equal(got, []int{1, 2, 3, 4}) {
		t.Errorf("FirstN([1 2 3 4], 10) = %v, want [1 2 3 4]", got)
	}
	if got := FirstN([]int{1, 2, 3, 4}, 0); len(got) != 0 {
		t.Errorf("FirstN([1 2 3 4], 0) = %v, want []", got)
	}
	if got := FirstN([]int{}, 3); len(got) != 0 {
		t.Errorf("FirstN([], 3) = %v, want []", got)
	}

	original := []int{1, 2, 3, 4}
	result := FirstN(original, 2)
	if len(result) > 0 {
		result[0] = 99
	}
	if original[0] != 1 {
		t.Errorf("after changing FirstN's result, the original changed too: %v (the result must be an independent copy)", original)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	input := []int{1, 2, 1, 3, 2}
	if got := RemoveDuplicates(input); !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("RemoveDuplicates([1 2 1 3 2]) = %v, want [1 2 3]", got)
	}
	if !slices.Equal(input, []int{1, 2, 1, 3, 2}) {
		t.Errorf("RemoveDuplicates changed its input: it is now %v, want [1 2 1 3 2]", input)
	}
	if got := RemoveDuplicates([]int{5, 5, 5}); !slices.Equal(got, []int{5}) {
		t.Errorf("RemoveDuplicates([5 5 5]) = %v, want [5]", got)
	}
	if got := RemoveDuplicates([]int{4, 3, 2}); !slices.Equal(got, []int{4, 3, 2}) {
		t.Errorf("RemoveDuplicates([4 3 2]) = %v, want [4 3 2]", got)
	}
	if got := RemoveDuplicates([]int{}); len(got) != 0 {
		t.Errorf("RemoveDuplicates([]) = %v, want []", got)
	}
}
