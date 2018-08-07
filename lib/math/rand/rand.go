package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Int returns a non-negative pseudo-random int from the default Source.
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	fmt.Println()

	// Intn returns, as an int, a non-negative pseudo-random number in [0,n)
	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d / ", r)
	}
	fmt.Println()

	// If Seed is not called, the generator behaves as if seeded by Seed(1).
	// Seed values that have the same remainder(余数) when divided by 2^31-1 generate the same pseudo-random sequence.
	// Seed, unlike the Rand.Seed method, is safe for concurrent use.
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}
}

/* Output:
134020434 / 1597969999 / 1721070109 / 2068675587 / 1237770961 / 220031192 / 2031484958 / 583324308 / 958990240 / 413002649 / 6 / 7 / 2 / 1 / 0 /
22.84 / 10.12 / 44.32 / 58.58 / 15.49 / 12.23 / 30.16 / 88.48 / 34.26 / 27.18 /
*/
