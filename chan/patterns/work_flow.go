package patterns

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/big"
	"sync"
)

// data generation/collecting/loading.
// A data producer may close the output
// stream channel at any time to end data generating
func RandomGenerator() <-chan uint64 {
	c := make(chan uint64)
	go func() {
		rnds := make([]byte, 8)
		for {
			_, err := rand.Read(rnds)
			if err != nil {
				close(c)
			}
			c <- binary.BigEndian.Uint64(rnds)
		}
	}()
	return c
}

// data aggregation/division
// A better implementation should consider whether or not
// an input stream has been closed.
func Aggregator(inputs ...<-chan uint64) <-chan uint64 {
	output := make(chan uint64)

	var wg sync.WaitGroup
	for _, in := range inputs {
		wg.Add(1)
		in := in // this line is essential
		go func() {
			// 负责处理一个数据流
			for {
				x, ok := <-in
				if ok {
					output <- x
				} else {
					wg.Done()
				}
			}
		}()
	}
	go func() {
		wg.Wait()     // 确保所有的输入数据流已关闭，且数据已经输出到output中
		close(output) // 关闭输出数据流
	}()
	return output
}

// data composition/decomposition.
// 将两个输入管道的数据，按照一定的规则，计算，输出到输出管道
// 如何优雅的关闭 output 管道? 协程如何退出？
func Composor(inA, inB <-chan uint64) <-chan uint64 {
	output := make(chan uint64)
	go func() {
		for {
			a1, b, a2 := <-inA, <-inB, <-inA // 每次：从A中读取两个数据，从B中读取一个数据
			output <- a1 ^ b&a2
		}
	}()
	return output
}

// data duplication/proliferation.
// 复制一份数据流，输出到两个管道
// 注意没有复用之前的管道，之前的管道是输入管道
func Duplicator(in <-chan uint64) (<-chan uint64, <-chan uint64) {
	outA, outB := make(chan uint64), make(chan uint64)
	go func() {
		for {
			x := <-in
			outA <- x
			outB <- x
		}
	}()
	return outA, outB
}

// data calculation/analysis.
// 读取输入管道的数据，按照一定的规则，计算，输出到输出管道
// 这个协程永远不会退出了！
// output如果不是自己的，那么也不应该负责关闭output管道。假如管道关闭以后，再往管道写数据，会导致panic ！
func Calculator(in <-chan uint64, out chan uint64) <-chan uint64 {
	if out == nil {
		out = make(chan uint64)
	}
	go func() {
		for {
			x := <-in
			out <- ^x
		}
	}()
	return out
}

// data validation/filtering.
func Filter(input <-chan uint64, output chan uint64) <-chan uint64 {
	if output == nil {
		output = make(chan uint64)
	}
	go func() {
		bigInt := big.NewInt(0)
		for {
			x := <-input
			bigInt.SetUint64(x)
			if bigInt.ProbablyPrime(1) {
				output <- x
			}
		}
	}()
	return output
}

// data serving/saving.
func Printer(input <-chan uint64) {
	for {
		x, ok := <-input
		if ok {
			fmt.Println(x)
		} else {
			return
		}
	}
}

func workFlow() {
	// 命名上都是以er或or结尾
	filterA := Filter(RandomGenerator(), nil)
	filterB := Filter(RandomGenerator(), nil)
	filterC := Filter(RandomGenerator(), nil)
	filter := Aggregator(filterA, filterB, filterC)

	// 同一个filter作为输出给到了两个 Calculator 手里, 扩展了 Calculator 的能力
	calculatorA := Calculator(filter, nil)
	calculatorB := Calculator(filter, nil)

	calculator := Aggregator(calculatorA, calculatorB)

	Printer(calculator)
}
