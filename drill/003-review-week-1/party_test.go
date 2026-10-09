package party

import "testing"

func TestTotalSlices(t *testing.T) {
	if got := TotalSlices(1); got != 8 {
		t.Errorf("TotalSlices(1) = %v, want 8", got)
	}
	if got := TotalSlices(3); got != 24 {
		t.Errorf("TotalSlices(3) = %v, want 24", got)
	}
}

func TestSlicesEach(t *testing.T) {
	if got := SlicesEach(2, 4); got != 4 {
		t.Errorf("SlicesEach(2, 4) = %v, want 4", got)
	}
	if got := SlicesEach(2, 5); got != 3 {
		t.Errorf("SlicesEach(2, 5) = %v, want 3", got)
	}
	if got := SlicesEach(1, 3); got != 2 {
		t.Errorf("SlicesEach(1, 3) = %v, want 2", got)
	}
}

func TestCostEach(t *testing.T) {
	if got := CostEach(2, 45.0, 4); got != 22.5 {
		t.Errorf("CostEach(2, 45.0, 4) = %v, want 22.5", got)
	}
	if got := CostEach(3, 40.0, 5); got != 24 {
		t.Errorf("CostEach(3, 40.0, 5) = %v, want 24", got)
	}
}

func TestSummary(t *testing.T) {
	want := "5 people, 16 slices, 3 slices each"
	if got := Summary(2, 5); got != want {
		t.Errorf("Summary(2, 5) = %q, want %q", got, want)
	}

	want = "4 people, 24 slices, 6 slices each"
	if got := Summary(3, 4); got != want {
		t.Errorf("Summary(3, 4) = %q, want %q", got, want)
	}
}
