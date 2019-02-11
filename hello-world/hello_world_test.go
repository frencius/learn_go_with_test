package main

import "testing"

func TestHelloWordWithName(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("saying hello world", func(t *testing.T) {
		got := HelloWord()
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to peole", func(t *testing.T) {
		got := HelloWordWithName("Leo")
		want := "Hello, Leo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("name is empty", func(t *testing.T) {
		got := HelloWordWithName("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello with another language", func(t *testing.T) {
		got := HelloWorldWithOtherLanguage("Spanish", "Leo")
		want := "Hola, Leo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := HelloWorldWithOtherLanguage("French", "Leo")
		want := "Bonjour, Leo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello with spain language with name empty", func(t *testing.T) {
		got := HelloWorldWithOtherLanguage("Spanish", "")
		want := "Hola, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello with spain language with name empty", func(t *testing.T) {
		got := HelloWorldWithOtherLanguage("French", "")
		want := "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello with another language with language empty", func(t *testing.T) {
		got := HelloWorldWithOtherLanguage("", "Leo")
		want := "Hello, Leo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello with another language with language and name empty", func(t *testing.T) {
		got := HelloWorldWithOtherLanguage("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
