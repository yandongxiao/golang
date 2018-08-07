package main

func main() {
	a := make(chan int)
	// deadlock
	if v, more := <-a; more {
		println(v)
	} else {
		println("---")
	}

}
