package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/big"
	"sync"
)

// data generation/collecting/loading.
// A data producer may close the output stream channel at any time to end data generating
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
// A better implementation should consider whether or not an input stream has been closed.
func Aggregator(inputs ...<-chan uint64) <-chan uint64 {
	output := make(chan uint64)
	var wg sync.WaitGroup
	for _, in := range inputs {
		wg.Add(1)
		in := in // this line is essential
		go func() {
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
		wg.Wait()
		close(output)
	}()
	return output
}

// data composition/decomposition.
func Composor(inA, inB <-chan uint64) <-chan uint64 {
	output := make(chan uint64)
	go func() {
		for {
			a1, b, a2 := <-inA, <-inB, <-inA
			output <- a1 ^ b&a2
		}
	}()
	return output
}

// data duplication/proliferation.
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

/*
func main() {
	Printer(Filter(Calculator(RandomGenerator(), nil), nil))
}
*/

func main() {
	filterA := Filter(RandomGenerator(), nil)
	filterB := Filter(RandomGenerator(), nil)
	filterC := Filter(RandomGenerator(), nil)
	filter := Aggregator(filterA, filterB, filterC)

	calculatorA := Calculator(filter, nil)
	calculatorB := Calculator(filter, nil)
	calculator := Aggregator(calculatorA, calculatorB)

	Printer(calculator)
}
