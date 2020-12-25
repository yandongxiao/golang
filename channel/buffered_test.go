// 注意：Design your algorithm in the first place with unbuffered channels.
// 注意：Only introduce buffering when performance is problematic.
package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestBufferedChannel(t *testing.T) {
	chMsg := make(chan string, 2)
	chMsg <- "hello"
	chMsg <- "world"
	close(chMsg)

	assert.Equal(t, <-chMsg, "hello")
	assert.Equal(t, <-chMsg, "world")
	assert.Equal(t, <-chMsg, "")
}
