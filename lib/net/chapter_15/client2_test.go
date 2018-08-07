package test

import (
	"errors"
	"io"
	"net"
	"testing"
)

func TestDialFunction(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:9527")
	if err != nil {
		t.Error(err)
		return
	}

	n, err := conn.Write([]byte("helLo"))
	if err != nil {
		t.Error(err)
		return
	}
	if n != 5 {
		//此种情况下我们需要继续写剩余的数据，编程模型立马变得复杂起来
		//同理，在Read时也可能会出现只读取一部分数据的可能性
		t.Error(errors.New("只写入部分数据"))
		return
	}

	for {
		buf := make([]byte, 1)
		n, err = conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				println("end of file")
				break
			}
			t.Error(err)
			return
		}
		if n != cap(buf) {
			t.Error(errors.New("只读取了部分数据"))
			return
		}
		println(string(buf))
	}
}
