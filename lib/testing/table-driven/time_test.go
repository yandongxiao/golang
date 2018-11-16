package main

import (
	"fmt"
	"testing"
	"time"
)

// a common way of writing tests in Go.
// This approach, commonly referred to as table-driven tests. makes it straightforward to add more test cases
func TestTime(t *testing.T) {
	testCases := []struct {
		gmt  string
		loc  string
		want string
	}{
		{"12:31", "Europe/Zuri", "13:31"},     // incorrect location name
		{"12:31", "America/New_York", "7:31"}, // should be 07:31
		{"08:08", "Australia/Sydney", "18:08"},
	}
	for _, tc := range testCases {
		loc, err := time.LoadLocation(tc.loc)
		if err != nil {
			t.Fatalf("could not load location %q", tc.loc)
		}
		gmt, _ := time.Parse("15:04", tc.gmt)
		if got := gmt.In(loc).Format("15:04"); got != tc.want {
			t.Errorf("In(%s, %s) = %s; want %s", tc.gmt, tc.loc, got, tc.want)
		}
	}
}

// Go 1.7 also introduces a Run method for creating subtests.
// Fatal and its siblings causes a subtest to be skipped but not its parent or subsequent subtests.
// Both subtests and sub-benchmarks can be singled out on the command line using the -run or -bench flag.
// Both flags take a slash-separated list of regular expressions that match the corresponding parts of the full name of the subtest or sub-benchmark.
// go test -run=TestTime/"in Europe"
// NOTE: Subtests and sub-benchmarks can be used to manage common setup and tear-down code
func TestTime(t *testing.T) {
	testCases := []struct {
		gmt  string
		loc  string
		want string
	}{
		{"12:31", "Europe/Zuri", "13:31"},
		{"12:31", "America/New_York", "7:31"},
		{"08:08", "Australia/Sydney", "18:08"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s in %s", tc.gmt, tc.loc), func(t *testing.T) {
			t.Parallel()
			loc, err := time.LoadLocation(tc.loc)
			if err != nil {
				t.Fatal("could not load location")
			}
			gmt, _ := time.Parse("15:04", tc.gmt)
			if got := gmt.In(loc).Format("15:04"); got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

// Before Go 1.7 it was not possible to use the same table-driven approach for benchmarks.
// A benchmark tests the performance of an entire function, so iterating over benchmarks would just measure all of them as a single benchmark.
// A common workaround was to define separate top-level benchmarks that each call a common function with different parameters.
func benchmarkAppendFloat(b *testing.B, f float64, fmt byte, prec, bitSize int) {
	dst := make([]byte, 30)
	b.ResetTimer()             // Overkill here, but for illustrative purposes.
	for i := 0; i < b.N; i++ { // b.N 由Benchmack指定循环次数
		AppendFloat(dst[:0], f, fmt, prec, bitSize)
	}
}

func BenchmarkAppendFloatDecimal(b *testing.B) { benchmarkAppendFloat(b, 33909, 'g', -1, 64) }
func BenchmarkAppendFloat(b *testing.B)        { benchmarkAppendFloat(b, 339.7784, 'g', -1, 64) }
func BenchmarkAppendFloatExp(b *testing.B)     { benchmarkAppendFloat(b, -5.09e75, 'g', -1, 64) }
func BenchmarkAppendFloatNegExp(b *testing.B)  { benchmarkAppendFloat(b, -5.11e-95, 'g', -1, 64) }
func BenchmarkAppendFloatBig(b *testing.B) {
	benchmarkAppendFloat(b, 123456789123456789123456789, 'g', -1, 64)
}

// Each invocation of the Run method creates a separate benchmark.
func BenchmarkAppendFloat(b *testing.B) {
	benchmarks := []struct {
		name    string
		float   float64
		fmt     byte
		prec    int
		bitSize int
	}{
		{"Decimal", 33909, 'g', -1, 64},
		{"Float", 339.7784, 'g', -1, 64},
		{"Exp", -5.09e75, 'g', -1, 64},
		{"NegExp", -5.11e-95, 'g', -1, 64},
		{"Big", 123456789123456789123456789, 'g', -1, 64},
	}
	dst := make([]byte, 30)
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) { // b.Run方法
			for i := 0; i < b.N; i++ {
				AppendFloat(dst[:0], bm.float, bm.fmt, bm.prec, bm.bitSize)
			}
		})
	}
}
