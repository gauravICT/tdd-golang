package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Gaurav")
	want := "Hello, Gaurav!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
