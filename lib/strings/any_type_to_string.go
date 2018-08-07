package main

type Person struct {
	name string
}

func main() {
	person := Person{"jack"}
	println(string(person))
}
