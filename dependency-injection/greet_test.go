package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Frencius")

	got := buffer.String()
	want := "Hello, Frencius"

	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
}
