package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestBytesBuffer(t *testing.T) {
	var buff bytes.Buffer

	buff.WriteString("helloworld")
	buff.WriteString("helloworld")
	buff.WriteString("helloworld")
	buff.WriteString("你好")

	assert.Equal(t, "helloworldhelloworldhelloworld你好", buff.String())
}

// 这是更加推荐的方式，尤其是不断调用 builder.String() 方法的情况下
func TestStringsBuilder(t *testing.T) {
	var builder strings.Builder

	builder.WriteString("helloworld")
	builder.WriteString("helloworld")
	builder.WriteString("helloworld")
	builder.WriteString("你好")

	assert.Equal(t, "helloworldhelloworldhelloworld你好", builder.String())
}
