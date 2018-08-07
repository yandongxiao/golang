package ttt

import "testing"

/* all test fuction must not be failed */
func Benchmark1_1(b *testing.B) {
	b.Error("benchmark_error")
	b.Log("execute here")
}

func Benchmark1_2(b *testing.B) {
	b.Fatal("benchmark_error")
	b.Log("do not execute here")
}

func Benchmark1_3(b *testing.B) {
	b.Log("OK")
}

func Benchmark1_4(b *testing.B) {
	b.StopTimer() /* have executed some times */
	b.Log("OK")
}

/* not only two times */
func Benchmark1_5(b *testing.B) {
	b.N = 1
	b.Log(b.N)
	for i := 0; i < b.N; i++ {
		b.Log("OK1")
	}
	b.Log("OK2")
}

func Benchmark1_6(b *testing.B) {
	num := new(int) /* NOTE: not right */
	*num = 1
	b.N = *num
	b.ReportAllocs()
}
