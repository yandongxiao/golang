package main

// Type I is a defined blank interface type.
type I interface{}

func main() {
	// interface{} is an unnamed blank interface type.
	// Two unnamed interface types are identical if their method sets are identical.
	var i interface{}
}
