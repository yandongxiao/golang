package test_test

import (
	"fmt"
	"os"
	"testing"
)

// if a test file contains a function: like this
func TestMain(m *testing.M) {
	// then the generated test will call TestMain(m) instead of
	// running the tests directly. TestMain runs in the main
	// goroutine and can do whatever setup and teardown
	// is necessary around a call to m.Run
	fmt.Println("setup")
	s := m.Run()
	fmt.Println("tear down")
	// It should then call os.Exit with the result of m.Run.
	os.Exit(s)
}

func TestHello(t *testing.T) {
	t.Log("hello")
}

func ExampleWorld() {
	fmt.Println("world")
	// Output:
	// world
}
