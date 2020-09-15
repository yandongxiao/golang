package fib

import "testing"

// Fib computes the n'th number in the Fibonacci series.
func Fib(n int) int {
	switch n {
	case 0:
		return 0
    case 1:
        return 1
    case 2:
        return 1
    default:
		return Fib(n-1) + Fib(n-2)
	}
}

func BenchmarkFib20(b *testing.B) {
    // 除了每次Fib(20)的执行时间，磁性次数
    // 该指令给出每次Fib(20)的申请内存，申请次数
    b.ReportAllocs()

    //  boringAndExpensiveSetup()
    b.ResetTimer()   // 去掉前面代码对性能统计的影响
	for n := 0; n < b.N; n++ {
        // If you have some expensive setup logic per loop iteration
        // b.StopTimer()
        // complicatedSetup()
        // b.StartTimer()
		Fib(20) // run the Fib function b.N times
	}
}

func TestFib(t *testing.T) {
	fibs := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}
	for n, want := range fibs {
		got := Fib(n)
		if want != got {
			t.Errorf("Fib(%d): want %d, got %d", n, want, got)
		}
	}
}

