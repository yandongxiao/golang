package main

import "io"

// 只要两个接口的close方法签名相同
type ReadWriter interface {
	io.ReadCloser
	io.WriteCloser
}

type readWriter struct {
}

func (r readWriter) Read(p []byte) (n int, err error) {
	panic("implement me")
}

func (r readWriter) Close() error {
	panic("implement me")
}

func (r readWriter) Write(p []byte) (n int, err error) {
	panic("implement me")
}

var (
	_ ReadWriter = readWriter(nil)
)
