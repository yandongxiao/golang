package main

import "fmt"
import "io/ioutil"

func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	entries, err := ioutil.ReadDir("/tmp")
	errcheck(err)
	for i := range entries {
		fmt.Printf("%v ", entries[i].Name())
	}

	fmt.Println()
}
