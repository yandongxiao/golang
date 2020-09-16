package main

import "testing"

func TestByDependencyInjection(t *testing.T) {
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)

	event.Start()
}

func TestByWire(t *testing.T) {
	e := InitializeEvent()
	e.Start()
}
