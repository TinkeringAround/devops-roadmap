package main

import "testing"

func TestWallet(t *testing.T) {
	checkBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	checkNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didnt wanted one")
		}
	}

	checkError := func(t testing.TB, got, want error) {
		t.Helper()

		if got == nil {
			t.Fatal("wanting an error but got none")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		checkBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(8))

		checkNoError(t, err)
		checkBalance(t, wallet, Bitcoin(12))
	})

	t.Run("withdraw insufficient amount", func(t *testing.T) {
		initBalance := Bitcoin(20)
		wallet := Wallet{initBalance}
		err := wallet.Withdraw(Bitcoin(100))

		checkError(t, err, ErrInsufficientFunds)
		checkBalance(t, wallet, initBalance)
	})
}
