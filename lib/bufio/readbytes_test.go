package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
)

func TestReadBytes(t *testing.T) {
	str := "helloworld"
	reader := bufio.NewReader(strings.NewReader(str))
	data, err := reader.ReadBytes('\n')
	assert.NotNil(t, err)
	assert.Equal(t, string(data), "helloworld")
}
