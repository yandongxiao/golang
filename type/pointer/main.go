package main

type T int

func (T) M() {}

func main() {
	var x **T
	// calling method M with receiver x (type **T) requires explicit dereference
	x.M()
}
