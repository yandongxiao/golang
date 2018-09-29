package main

import "fmt"

// First of to remove an item from a slice you need to use built-in function append:
func useAppend() {
	daysOfWeek := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	daysOfWeek = append(daysOfWeek[:1], daysOfWeek[2:]...)
	fmt.Println(daysOfWeek)
}

type Person struct {
	Name   string
	Remove bool
}

var people = []Person{
	{"P0", false},
	{"P1", false},
	{"P2", false},
	{"P3", true},
}

// 竟然是成功的?
func useForRange() {

	for i, p := range people {
		if p.Remove {
			people = append(people[:i], people[i+1:]...)
		}
	}
	fmt.Println(people)
}

// correct
func useReverse() {
	for i := len(people) - 1; i >= 0; i-- {
		if people[i].Remove {
			people = append(people[:i], people[i+1:]...)
		}
	}
	fmt.Println(people) // [{P0 false} {P2 false}]
}

func main() {
	useForRange()
	useReverse()
}
