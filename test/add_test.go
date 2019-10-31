package test_test

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"

	"github.com/yandongxiao/golang-learning/test"
)

func TestAdd(t *testing.T) {
	a := 10
	b := 10
	c := test.Add(a, b)
	// printed only if the test fails or the -test.v flag is set.
	t.Log("helloworld", t.Name(), c)
}

func TestSkip(t *testing.T) {
	t.Log("hello")
	t.Skip("not implemented, run to here")
	t.Log("world")
}

// NOTE: Benchmarks are run sequentially.
// go test -bench=.	// 执行BenchmarkXxx
// To start tuning the Go program, we have to enable profiling.
// If the code used the Go testing package's benchmarking support,
// we could use gotest's standard -cpuprofile and -memprofile flags
func BenchmarkAdd(b *testing.B) {
	// If a benchmark needs some expensive setup before running, the timer may be reset
	b.ResetTimer()
	// The benchmark function must run the target code b.N times.
	for i := 0; i < b.N; i++ {
		c := test.Add(i, i)
		if c != 2*i {
			fmt.Println("helloworld")
		}
	}
}

//NOTE: go test -bench=. -cpu=2(指定并行数)
// If a benchmark needs to test performance in a parallel setting,
// it may use the RunParallel helper function
// such benchmarks are intended to be used with the go test -cpu flag
func BenchmarkTemplateParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}
