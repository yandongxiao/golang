package main

type bar struct {
}

func (b *bar) Error() string {
	return "bar"
}

// error is an interface, and *bar implement it
// rule: if err == nil; means the underlying type is nil, and the underlying value is nil.
// err = b; means the underlying type is not nil.
func foo() error {
	var b *bar

	// ...

	// NOTE: correct: return nil if b is nil.
	return b // b == nil
}

func main() {
	if foo() != nil { // the underlying type of interface error is not nil
		panic("foo")
	}
}
