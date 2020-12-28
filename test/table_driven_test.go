package test_test

import (
	"fmt"
	"testing"

	"github.com/yandongxiao/go/test"
)

// The Run methods of T and B allow defining subtests
// and sub-benchmarks, without having to define separate
// functions for each. This enables uses like table-driven
// benchmarks and creating hierarchical tests. It also
// provides a way to share common setup and tear-down code:
func TestAdd_suffix(t *testing.T) {
	cases := []struct {
		a   int
		b   int
		sum int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{4, 1, 5},
	}

	// Each subtest and sub-benchmark has a unique name
	// the combination of the name of the top-level test
	// and the sequence of names passed to Run, separated by slashes
	for i := range cases {
		i := i
		// It runs f in a separate goroutine and blocks until f returns
		// or calls t.Parallel to become a parallel test.
		t.Run(fmt.Sprintf("Add(%d,%d)=%d",
			cases[i].a, cases[i].b, cases[i].sum), func(t *testing.T) {
			// all tests are run in parallel with each other
			// The race detector kills the program if it exceeds 8192 concurrent
			// goroutines, so use care when running parallel tests with
			// the -race flag set.
			t.Parallel()
			sum := test.Add(cases[i].a, cases[i].b)
			if sum != cases[i].sum {
				t.Errorf("Add(%d, %d) == %d, want: %d",
					cases[i].a, cases[i].b, sum, cases[i].sum)
			}
		})
	}
}

func TestAdd_suffix2(t *testing.T) {
	cases := []struct {
		a   int
		b   int
		sum int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{4, 1, 5},
	}

	// 注意与TestAdd_suffix的区别，t.Parallel()会导致Run提前返回的
	// Run does not return until parallel subtests have completed,
	// providing a way to clean up after a group of parallel tests
	t.Run("group", func(t *testing.T) {
		for i := range cases {
			c := cases[i]
			t.Run(fmt.Sprintf("Add(%d,%d)=%d",
				c.a, c.b, c.sum), func(t *testing.T) {
				sum := test.Add(c.a, c.b)
				if sum != cases[i].sum {
					t.Errorf("Add(%d, %d) == %d, want: %d",
						c.a, c.b, sum, c.sum)
				}
			})
		}
	})
}
