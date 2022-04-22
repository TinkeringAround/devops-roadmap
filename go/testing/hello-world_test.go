package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		you := "Thomas"
		got := Hello(you, "English")
		want := "Hello, " + you

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to World when you is empty string", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in french", func(t *testing.T) {
		got := Hello("Francoise", "French")
		want := "Bonjour, Francoise"

		assertCorrectMessage(t, got, want)
	})
}
