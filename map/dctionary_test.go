package _map

import (
	"errors"
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Errorf("should have returned an error")
		}
		assertError(t, err, ErrNotFound)
	})
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}

	fmt.Println()
}
