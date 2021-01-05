package main

import "testing"

func TestReallySlowGreeting(t *testing.T) {
	expectedGreeting := "Code.education Rocks!"

	receivedGreeting := reallySlowGreeting()

	if receivedGreeting != expectedGreeting {
		t.Errorf("Greeting should be %s, but was %s", expectedGreeting, receivedGreeting)
	}
}