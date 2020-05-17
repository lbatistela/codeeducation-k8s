package main

import "testing"

func TestGreeting(t *testing.T) {
	actual := greeting("This is a test")
	expected := "<b>This is a test</b>"
	if actual != expected {
		t.Errorf("Expected [%v], got [%v]", expected, actual)
	}
}