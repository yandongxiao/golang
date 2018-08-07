package main

import (
	"fmt"
	"sort"
)

func main() {
	p1 := Person{"alice", 12}
	p2 := Person{"bob", 13}
	p3 := Person{"tom", 1}
	persons := []*Person{&p1, &p2, &p3}
	sort.Sort(PersonSlice(persons))

	// f(i) == true implies f(i+1) == true
	age := 1
	idx := sort.Search(len(persons), func(i int) bool {
		fmt.Println(persons[i])
		return persons[i].age >= age
	})
	println(idx)
}

// func Search(n int, f func(int) bool) int
// Search uses binary search to find and return the smallest index i in [0, n) at which f(i) is true,
// assuming that on the range [0, n), f(i) == true implies f(i+1) == true.
// Search returns the first true index.
// If there is no such index, Search returns n.
type Person struct {
	name string
	age  int
}

type PersonSlice []*Person

func (ps PersonSlice) Len() int {
	return len(ps)
}

func (ps PersonSlice) Less(i, j int) bool {
	if ps[i].age < ps[j].age {
		return true
	}
	return false
}

func (ps PersonSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
