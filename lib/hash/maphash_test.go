package main

import (
	"fmt"
	"hash/maphash"
)

func ExampleMapHash() {
	// Package maphash provides hash functions on byte sequences.
	// 它和map不是互斥的关系

	// The zero Hash value is valid and ready to use; setting an
	// initial seed is not necessary.
	// A zero Hash chooses a random seed for itself during the first
	// call to a Reset, Write, Seed, Sum64, or Seed method.
	var h maphash.Hash

	// Add a string to the hash, and print the current hash value.
	h.WriteString("hello, ")
	fmt.Printf("%#x\n", h.Sum64()) // 这个就是hash值

	// Append additional data (in the form of a byte array).
	h.Write([]byte{'w', 'o', 'r', 'l', 'd'})
	fmt.Printf("%#x\n", h.Sum64())

	// Reset discards all data previously added to the Hash, without
	// changing its seed.
	h.Reset()

	// Use SetSeed to create a new Hash h2 which will behave
	// identically to h.
	var h2 maphash.Hash
	h2.SetSeed(h.Seed())

	h.WriteString("same")
	h2.WriteString("same")
	fmt.Printf("%#x == %#x\n", h.Sum64(), h2.Sum64())

	// Output:
	//
}
