package main

import "testing"

func assertBalance(t testing.TB, w Wallet, want Bitcoin){
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error){
	t.Helper()
	if got == nil {
		t.Errorf("Did not get an error but wonted one")
	}
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error){
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}


func TestWallet(t *testing.T)  {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		got :=wallet.Withdraw(Bitcoin(10))

		assertNoError(t, got)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		w := Wallet{startingBalance}
		err := w.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, w, startingBalance)
	})
}