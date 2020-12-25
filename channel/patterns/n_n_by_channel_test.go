package patterns

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type T = struct{}

// ready 用来接收信号，开始工作
// done 用来发送信号，结束工作
func worker(id int, ready <-chan T, done chan<- T, dataC chan<- int) {
	<-ready // block here and wait a notification

	log.Print("Worker#", id, " starts.")
	// Simulate a workload.
	time.Sleep(time.Millisecond * 100 * time.Duration(id+1))
	dataC <- id
	log.Print("Worker#", id, " job done.")

	// Notify the main goroutine (N-to-1)
	done <- T{}
}

func TestNToNByChannel(t *testing.T) {
	// ready 用来通知worker开始工作
	// done 用来接收worker的信息，表示工作已完成
	// dataC 用来接收worker的结果，必须是 buffered channel
	ready, done, dataC := make(chan T), make(chan T), make(chan int, 3)
	go worker(0, ready, done, dataC)
	go worker(1, ready, done, dataC)
	go worker(2, ready, done, dataC)

	// 1-to-N notifications.
	// close(ready) // broadcast notifications, 用这个来替代下面三行代码
	ready <- T{}
	ready <- T{}
	ready <- T{}

	// Being N-to-1 notified.
	<-done
	<-done
	<-done

	assert.Equal(t, <-dataC+<-dataC+<-dataC, 3)
}
