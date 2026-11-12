package methods

import "testing"

func TestArea(t *testing.T) {
	r := Rectangle{Width: 3, Height: 4}
	if got := r.Area(); got != 12 {
		t.Errorf("%+v.Area() = %d, want 12", r, got)
	}
	line := Rectangle{Width: 0, Height: 7}
	if got := line.Area(); got != 0 {
		t.Errorf("%+v.Area() = %d, want 0", line, got)
	}
}

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

func TestDeposit(t *testing.T) {
	acc := Account{Owner: "Ana", Balance: 100}
	acc.Deposit(50)
	if acc.Balance != 150 {
		t.Fatalf("after Deposit(50) on 100, Balance = %d, want 150", acc.Balance)
	}
	acc.Deposit(-20)
	if acc.Balance != 150 {
		t.Errorf("after Deposit(-20) on 150, Balance = %d, want 150", acc.Balance)
	}
	acc.Deposit(0)
	if acc.Balance != 150 {
		t.Errorf("after Deposit(0) on 150, Balance = %d, want 150", acc.Balance)
	}
}

func TestWithdraw(t *testing.T) {
	acc := Account{Owner: "Ana", Balance: 100}
	if got := acc.Withdraw(30); got != true || acc.Balance != 70 {
		t.Fatalf("Withdraw(30) on 100 = %v with Balance %d, want true with Balance 70", got, acc.Balance)
	}
	if got := acc.Withdraw(500); got != false || acc.Balance != 70 {
		t.Errorf("Withdraw(500) on 70 = %v with Balance %d, want false with Balance 70", got, acc.Balance)
	}
	acc.Balance = 70
	if got := acc.Withdraw(-5); got != false || acc.Balance != 70 {
		t.Errorf("Withdraw(-5) on 70 = %v with Balance %d, want false with Balance 70", got, acc.Balance)
	}
	acc.Balance = 70
	if got := acc.Withdraw(70); got != true || acc.Balance != 0 {
		t.Errorf("Withdraw(70) on 70 = %v with Balance %d, want true with Balance 0", got, acc.Balance)
	}
}
