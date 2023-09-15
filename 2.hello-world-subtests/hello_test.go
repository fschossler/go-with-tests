package main

import "testing"

// This function runs two subtests for testing the two possible variables of the Hello function(with a name and empty).
func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "") // This second empty string is a "gambiarra" in my opinion but the author is ok with this
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	// This function is needed to tell the compiler that is a helper function
	// in this way the output of error show the line with error in the test not in
	// this function.
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
