package main

import (
	"fmt"
	"time"
)

func work(workerId int, jobs <-chan int, results chan<- int) {
	for jobId := range jobs {
		fmt.Println("workerId: ", workerId, " receive job", jobId)
		time.Sleep(time.Second)
		fmt.Println("workerId: ", workerId, " job", jobId, "done")
		results <- jobId * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 0; i < 1; i++ {
		go work(i, jobs, results)
	}

	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs)

	println("dispatch jobs done")
	for i := 0; i < 10; i++ {
		println(<-results)
	}

}
