package main

func main() {
	// A common idiom used to let the main program block indefinitely while other goroutines run is to place select {} as the last statement in a main function.}
	// 但是如果没有其它协程，会发生deadlock
	// The default clause is optional; fall through behavior, like in the normal switch, is not permitted.
	// If there are no cases, the select blocks execution forever.
	// 一般与for一起使用.  A select is terminated when a break or return is executed in one of its cases. (这句话有问题吧)
	select {}
}
