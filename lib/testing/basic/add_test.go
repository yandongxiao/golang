// These test-programs must be within the same package and the files must have names of the form *_test.go, so the test code is separated from the actual code of the package.
package main

import "testing"

func TestAdd(t *testing.T) {
	if add(1, 2) == 3 {
		t.Logf("add test success")
	}
}

// The command go test –.bench=. -run=none executes all these functions
// they will call the functions in the code a very large number of times N (e.g. N = 1000000),
// show this N and the average execution time of the functions in ns (ns/op)
func BenchmarkReverse(b *testing.B) {
	b.Log("hellowolrd")
}

func TestAdd_2(t *testing.T) {
	failure(t)
}

// Helper marks the calling function as a test helper function.
func failure(t *testing.T) {
	t.Helper() // This call silences this function in error reports.
	t.Fatal("failure")
}
