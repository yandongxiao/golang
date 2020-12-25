// NOTE: all types can be used in a channel.(包括 channel)
package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestChannelInChannel(t *testing.T) {
	resp := make(chan int)
	req := make(chan chan int)

	go func() {
		resp := <-req
		resp <- 1
		close(resp)
	}()

	req <- resp
	for v := range resp {
		assert.Equal(t, v, 1)
	}
}

func TestInterfaceInChannel(t *testing.T) {
	c := make(chan interface{})
	go func() {
		c <- 3
	}()

	assert.Equal(t, <-c, 3)
}
