package main

import "testing"

func TestDictionary_SearchEntry(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	checkString := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	}

	checkError := func(t testing.TB, got, want error) {
		t.Helper()

		if got != want {
			t.Errorf("got error %q want %q", got, want)
		}
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.SearchEntry("test")
		want := "this is just a test"

		checkString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.SearchEntry("unknown-word")

		checkError(t, got, ErrWordNotFound)
	})
}

func TestDictionary_AddEntry(t *testing.T) {
	dictionary := Dictionary{"existing": "existing definition"}

	checkString := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	}

	checkError := func(t testing.TB, got, want error) {
		t.Helper()

		if got != want {
			t.Errorf("got error %q want %q", got, want)
		}
	}

	t.Run("should add new word for non-existing key to dictionary", func(t *testing.T) {
		var (
			key        = "not-existing"
			definition = "this is just a test"
		)
		_ = dictionary.AddEntry(key, definition)

		got, err := dictionary.SearchEntry(key)

		checkError(t, err, nil)
		checkString(t, got, definition)
	})

	t.Run("should throw error for adding entry of an existing word", func(t *testing.T) {
		err := dictionary.AddEntry("existing", "overwrite existing key")

		checkError(t, err, ErrWordExists)
	})
}

func TestDictionary_UpdateEntry(t *testing.T) {
	dictionary := Dictionary{"existing": "existing definition"}

	checkString := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	}

	checkError := func(t testing.TB, got, want error) {
		t.Helper()

		if got != want {
			t.Errorf("got error %q want %q", got, want)
		}
	}

	t.Run("should update definition for existing key", func(t *testing.T) {
		var (
			key        = "existing"
			definition = "new definition"
		)
		err := dictionary.UpdateEntry(key, definition)
		got, _ := dictionary.SearchEntry(key)

		checkError(t, err, nil)
		checkString(t, got, definition)
	})

	t.Run("should throw error when updating non-existing key", func(t *testing.T) {
		err := dictionary.UpdateEntry("not-existing", "overwrite existing key")

		checkError(t, err, ErrWordNotExisting)
	})
}

func TestDictionary_DeleteEntry(t *testing.T) {
	dictionary := Dictionary{"existing": "existing definition"}

	checkError := func(t testing.TB, got, want error) {
		t.Helper()

		if got != want {
			t.Errorf("got error %q want %q", got, want)
		}
	}

	t.Run("known word", func(t *testing.T) {
		dictionary.DeleteEntry("existing")

		_, err := dictionary.SearchEntry("existing")
		checkError(t, err, ErrWordNotFound)
	})
}
