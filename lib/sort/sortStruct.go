package main

import (
	"fmt"
	"sort"
)

// srot.Interface
//  A type, typically a collection, that satisfies sort.Interface can be sorted by the routines in this package.
//  The methods require that the elements of the collection be enumerated by an integer index.
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

func main() {
	p1 := Person{"alice", 12}
	p2 := Person{"bob", 13}
	p3 := Person{"tom", 1}
	persons := []*Person{&p1, &p2, &p3}
	sort.Sort(PersonSlice(persons))

	for i := range persons {
		p := persons[i]
		fmt.Println(*p)
	}
}
