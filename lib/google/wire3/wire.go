// +build wireinject

package main

import "github.com/google/wire"

func InitializeEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil // 和 NewEvent 的返回值类型保持一致
}
