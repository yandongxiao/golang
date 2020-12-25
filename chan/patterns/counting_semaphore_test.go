package patterns

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

// 酒吧
type Bar chan Seat

// 座位
type Seat int

var money = make(chan int)

func (bar Bar) ServeCustomer(customerId int) {
	// log.Print("customer#", customerId, " enters the bar")

	seat := <-bar // need a seat to drink
	// log.Print("++ customer#", customerId, " drinks at seat#", seat)

	money <- 2 // 假设每人消费两元

	// log.Print("-- customer#", customerId, " frees seat#", seat)
	bar <- seat // free seat and leave the bar
}

func TestCountingSemaphore(t *testing.T) {
	// 假设酒吧有十把椅子
	rand.Seed(time.Now().UnixNano())
	bar24x7 := make(Bar, 10)

	// 将十把椅子摆到酒吧内，准备迎客
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		// None of the sends will block.
		bar24x7 <- Seat(seatId)
	}

	// 什么时候结束营业，收银员说了算
	done := make(chan struct{})
	sum := 0
	go func() {
		for v := range money {
			sum += v
		}
		close(done)
	}()

	// 相当于请求在不断的到来，每个请求对应一个协程。通过Counting Semaphore机制，限制处理请求的最大并发数。
	var wg sync.WaitGroup
	customers := 10000
	for customerId := 0; customerId < customers; customerId++ {
		wg.Add(1)
		// 这里有点问题，10000名顾客都进到了酒吧里，等待位置
		go func(id int) {
			defer wg.Done()
			bar24x7.ServeCustomer(id)
		}(customerId)
	}
	wg.Wait()
	close(money)
	<-done

	assert.Equal(t, sum, customers*2)
}
