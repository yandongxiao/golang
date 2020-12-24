package buildin

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestArrayLen(t *testing.T) {
	assert.Equal(t, len([3]int{}), 3)
	assert.Equal(t, len(&[3]int{}), 3) // cap() 作用于指针时，同理
	var p *[3]int
	assert.Equal(t, len(p), 3)
}

func TestChanLen(t *testing.T) {
	ch := make(chan int, 1)
	assert.Equal(t, len(ch), 0)

	ch <- 1
	assert.Equal(t, len(ch), 1)

	close(ch)
	assert.Equal(t, len(ch), 1)

	<-ch
	assert.Equal(t, len(ch), 0)
	fmt.Println(len(ch))
}
