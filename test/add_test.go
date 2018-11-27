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

// Benchmarks are run sequentially.
// The benchmark function must run the target code b.N times.
// go help testflag: go test -bench=.
func BenchmarkAdd(b *testing.B) {
	// If a benchmark needs some expensive setup before running, the timer may be reset
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := test.Add(i, i)
		if c != 2*i {
			// TODO
			fmt.Println("helloworld")
		}
	}
}

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
