package main

import "testing"

// func TestHello(t *testing.T) {
// 	got := Hello("Gaurav")
// 	want := "Hello, Gaurav!"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// Subtests :
// func TestHello(t *testing.T) {

// 	t.Run("saying hello to people", func(t *testing.T) {
// 		got := Hello("Chris")
// 		want := "Hello, Chris"

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})

// 	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
// 		got := Hello("")
// 		want := "Hello, World"

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})

// }

// Refactor :
func TestHello(t *testing.T) {
	assertCorrectMsg := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Gaurav", "spanish")
		want := "Hello, Gaurav"
		assertCorrectMsg(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMsg(t, got, want)
	})
}
