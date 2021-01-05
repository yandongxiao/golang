package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
)

func ExampleReadSlice(t *testing.T) {
	reader := bufio.NewReader(
		strings.NewReader("studygolang.com. \nIt is the home of gophers"))

	// line 的内容包含 \n
	// 注意： ReadSlice fails with error ErrBufferFull if the buffer fills without a delim.
	line, err := reader.ReadSlice('\n')
	assert.Nil(t, err)
	origin := string(line)

	_, _ = reader.ReadSlice('\n')
	assert.NotEqual(t, origin, string(line))
}
