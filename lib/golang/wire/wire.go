// +build wireinject

package main

import "github.com/google/wire"

// InitializeEvent call injector
// Dependency Injection is passing dependency to other objects or framework( dependency injector).
// the injector's purpose is to provide information about which providers to use to construct an Event
func InitializeEvent() Event {
	// NewEvent, NewGreeter, NewMessage called initializers or providers
	// providers are functions which provide a particular type
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
