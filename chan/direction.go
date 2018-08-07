package main

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// only send message
	go func(sender chan<- string) {
		sender <- "helloworld"
	}(ch1)

	go func(receiver <-chan string, sender chan<- string) {
		msg := <-receiver
		sender <- msg
	}(ch1, ch2)

	println(<-ch2)
}
