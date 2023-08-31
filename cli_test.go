package main

import "testing"

func TestCleanInputLower(t *testing.T) {
	input := "Hello"
	expected := "hello"
	actual := clean_input("HELLO")
	if expected != actual {
		t.Errorf("clean_input(%q) == %q, expected %q", input, actual, expected)
	}
}
