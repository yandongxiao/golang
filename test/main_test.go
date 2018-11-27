package test_test

import (
	"fmt"
	"os"
	"testing"
)

// if a test file contains a function: like this
func TestMain(m *testing.M) {
	// the generated test will call TestMain(m) instead of running
	// the tests directly. TestMain runs in the main goroutine and
	// can do whatever setup and teardown is necessary around a call to m.Run
	fmt.Println("setup")
	s := m.Run()
	fmt.Println("tear down")
	os.Exit(s)
}
