package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Chris", emptyString)
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to ", func(t *testing.T){
		got := Hello(emptyString, emptyString)
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish ", func(t *testing.T){
		got := Hello("Elodie", spanishLanguage)
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in French ", func(t *testing.T){
		got := Hello("Remi", frenchLanguage)
		want := "Bonjour, Remi"

		assertCorrectMessage(t, got, want)
	})

}
