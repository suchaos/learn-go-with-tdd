package hello_world

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		result := Hello("tom", "")
		excepted := "hello, tom"

		assertCorrectMessage(t, result, excepted)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		result := Hello("", "")
		excepted := "hello, world"

		assertCorrectMessage(t, result, excepted)
	})

	t.Run("in Spanish", func(t *testing.T) {
		result := Hello("Elodie", spanish)
		want := "hola, Elodie"

		assertCorrectMessage(t, result, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Lauren", french)
		want := "bonjour, Lauren"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, result string, excepted string) {
	t.Helper()

	if result != excepted {
		t.Errorf("excepted %s, but got %s", excepted, result)
	}
}
