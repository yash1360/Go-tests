package main

import (
	"testing"
)

func TestIntro(t *testing.T) {
	t.Run("saying hello user when name sting is empty", func(t *testing.T) {
		got := Hello("")
		want := "Hello, User"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello user_name when name sting is user_name", func(t *testing.T) {
		got := Hello("Yash")
		want := "Hello, Yash"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
