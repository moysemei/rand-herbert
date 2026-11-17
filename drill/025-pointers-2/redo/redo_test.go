package redo

import "testing"

func TestFits(t *testing.T) {
	small := Rectangle{Width: 2, Height: 5}
	if got := small.Fits(Rectangle{Width: 3, Height: 6}); got != true {
		t.Errorf("{2 5}.Fits({3 6}) = %v, want true (standing)", got)
	}
	if got := small.Fits(Rectangle{Width: 6, Height: 3}); got != true {
		t.Errorf("{2 5}.Fits({6 3}) = %v, want true (turned: 5 along the width, 2 along the height)", got)
	}
	square := Rectangle{Width: 4, Height: 4}
	if got := square.Fits(Rectangle{Width: 3, Height: 6}); got != false {
		t.Errorf("{4 4}.Fits({3 6}) = %v, want false", got)
	}
	long := Rectangle{Width: 7, Height: 1}
	if got := long.Fits(Rectangle{Width: 6, Height: 3}); got != false {
		t.Errorf("{7 1}.Fits({6 3}) = %v, want false (7 is too long both ways)", got)
	}
	same := Rectangle{Width: 3, Height: 6}
	if got := same.Fits(Rectangle{Width: 3, Height: 6}); got != true {
		t.Errorf("{3 6}.Fits({3 6}) = %v, want true (touching the border counts)", got)
	}
}
