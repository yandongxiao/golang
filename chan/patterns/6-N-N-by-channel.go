package main

import "log"
import "time"

type T = struct{}

func worker(id int, ready <-chan T, done chan<- T) {
	<-ready // block here and wait a notification
	log.Print("Worker#", id, " starts.")
	// Simulate a workload.
	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("Worker#", id, " job done.")
	// Notify the main goroutine (N-to-1),
	done <- T{}
}
func main() {
	log.SetFlags(0)
	ready, done := make(chan T), make(chan T)
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)
	// Simulate an initialization phase.
	time.Sleep(time.Second * 3 / 2)
	// 1-to-N notifications.
	// close(ready) // broadcast notifications, 用这个来替代下面三行代码
	ready <- T{}
	ready <- T{}
	ready <- T{}
	// Being N-to-1 notified.
	<-done
	<-done
	<-done
}
