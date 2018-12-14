package main

import "fmt"

var counter = 10

func ExampleRace() {
	go func() {
		counter++
	}()
	counter++
	fmt.Println(counter)
}
