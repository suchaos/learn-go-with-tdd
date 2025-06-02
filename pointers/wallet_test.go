package pointers

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		want := Bitcode(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcode(20)}
		err := wallet.Withdraw(Bitcode(10))

		assertNotError(t, err)
		assertBalance(t, wallet, Bitcode(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcode(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcode(100))

		assertError(t, err, ErrInsufficientFounds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertNotError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, err error, want error) {
	t.Helper()

	if err == nil {
		t.Fatal("expected an error but didn't get one")
	}

	if !errors.Is(err, want) {
		t.Errorf("got %q want %q", err, want)
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcode) {
	t.Helper()

	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
