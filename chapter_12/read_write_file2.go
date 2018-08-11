package main

import "io/ioutil"

func main() {
	data, err := ioutil.ReadFile("input.dat")
	if err != nil {
		println("read error")
	}

	ioutil.WriteFile("output.dat", data, 0644)
}
