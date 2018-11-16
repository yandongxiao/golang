// go tool cover -h
// go test -coverprofile=c.out
// go tool cover -html=c.out
package main

import "testing"

func TestFoo(t *testing.T) {
	foo()
}
