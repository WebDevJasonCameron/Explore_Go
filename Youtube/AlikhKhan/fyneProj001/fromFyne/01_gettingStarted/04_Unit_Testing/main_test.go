package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestGreeting(t *testing.T) {
	out, in := makeUI()

	if out.Text != "Hello World!" {
		t.Error("Incorrect initial greeting")
	}

	test.Type(in, "Andy")

	if out.Text != "Hello Andy!" {
		t.Error("Incorrect user greeting")
	}
}
