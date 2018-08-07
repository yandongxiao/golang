package main

func main() {
	chMsg := make(chan string)
	go func() {
		chMsg <- "hello"
		chMsg <- "world"
		close(chMsg) // 如果注释这条语句，将会出现死锁
	}()

	// This range iterates over each element as it’s received from queue.
	// 直到chan closed
	for msg := range chMsg {
		println(msg)
	}
}
