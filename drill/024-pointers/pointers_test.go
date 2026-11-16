package pointers

import (
	"slices"
	"testing"
)

func TestDouble(t *testing.T) {
	x := 5
	Double(&x)
	if x != 10 {
		t.Errorf("after Double(&x) with x = 5, x = %d, want 10", x)
	}
	y := -3
	Double(&y)
	if y != -6 {
		t.Errorf("after Double(&y) with y = -3, y = %d, want -6", y)
	}
}

func TestSwap(t *testing.T) {
	x, y := 1, 2
	Swap(&x, &y)
	if x != 2 || y != 1 {
		t.Errorf("after Swap(&x, &y) with x = 1, y = 2: x = %d, y = %d, want x = 2, y = 1", x, y)
	}
}

func TestAddBonus(t *testing.T) {
	accs := []Account{{Owner: "Ana", Balance: 100}, {Owner: "Bia", Balance: 50}}
	AddBonus(accs, 10)
	want := []Account{{Owner: "Ana", Balance: 110}, {Owner: "Bia", Balance: 60}}
	if !slices.Equal(accs, want) {
		t.Errorf("after AddBonus(accs, 10), accs = %+v, want %+v", accs, want)
	}
	AddBonus([]Account{}, 10) // must not crash on an empty slice
}

func TestFindByOwner(t *testing.T) {
	accs := []Account{{Owner: "Ana", Balance: 100}, {Owner: "Bia", Balance: 50}}
	p := FindByOwner(accs, "Bia")
	if p == nil {
		t.Fatalf("FindByOwner(accs, %q) = nil, want a pointer to Bia's account", "Bia")
	}
	if *p != (Account{Owner: "Bia", Balance: 50}) {
		t.Fatalf("FindByOwner(accs, %q) points to %+v, want {Owner:Bia Balance:50}", "Bia", *p)
	}
	p.Balance += 5
	if accs[1].Balance != 55 {
		t.Errorf("after p.Balance += 5, accs[1].Balance = %d, want 55 (p must point into the slice, not to a copy)", accs[1].Balance)
	}
	if got := FindByOwner(accs, "Caio"); got != nil {
		t.Errorf("FindByOwner(accs, %q) = %+v, want nil", "Caio", *got)
	}
	if got := FindByOwner([]Account{}, "Ana"); got != nil {
		t.Errorf("FindByOwner([], %q) = %+v, want nil", "Ana", *got)
	}
}
