package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func test1() {
	fReader, _ := os.Open("/tmp/data.txt")
	bufReader := bufio.NewReader(fReader)

	// Buffered returns the number of bytes that can be read from the current buffer.
	// 因为还没有进行任何读取操作，返回0
	fmt.Println(bufReader.Buffered())

	// The bytes are taken from at most one Read on the underlying Reader, hence n may be less than len(p)
	// 具体读取到多少数据、是否缓存数据，完全取决于底层Reader的行为。所以，不保证Read到10个字节
	// At EOF, the count will be zero and err will be io.EOF. 说明，如果读取部分数据后读到文件末尾，Read返回n, nil
	data := make([]byte, 10)
	fmt.Println(bufReader.Read(data))

	fmt.Println(bufReader.Peek(10))

	fmt.Println(bufReader.Buffered())
}

func test2() {
	fReader, _ := os.Open("/tmp/data.txt")
	bufReader := bufio.NewReader(fReader)
	// Peek returns the next n bytes without advancing the reader. 只有拥有缓存功能的Reader才可以提供这个功能
	fmt.Println(bufReader.Peek(10))
	time.Sleep(5 * time.Second) // 在此期间修改文件内容
	// The bytes stop being valid at the next read call.
	// 每次Peek操作也会触发底层Reader执行Read操作，返回最新的数据内容
	fmt.Println(bufReader.Peek(10))
}

func test3() {
	fReader, _ := os.Open("/tmp/data.txt")
	bufReader := bufio.NewReader(fReader)
	// If Discard skips fewer than n bytes, it also returns an error.
	// If 0 <= n <= b.Buffered(), Discard is guaranteed to succeed without reading from the underlying io.Reader.
	fmt.Println(bufReader.Discard(2))
}

func test4() {
	fReader, _ := os.Open("/tmp/data.txt")
	bufReader := bufio.NewReader(fReader)
	for {
		// 既然go已经缓存了读取的数据，所以下面的操作也是有效率的
		b, err := bufReader.ReadByte()
		if err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Printf("%c", b)
		}
	}
}

func test5() {
	fReader, _ := os.Open("/tmp/data.txt")
	bufReader := bufio.NewReader(fReader)
	for {
		if data, err := bufReader.ReadBytes(' '); err != nil {
			fmt.Printf("%v %v\n", string(data[:len(data)-1]), err)
			break
		} else {
			fmt.Printf("%v\n", string(data[:len(data)-1]))
		}
	}
}

func test6() {
	fReader, _ := os.Open("/tmp/data.txt")
	bufReader := bufio.NewReader(fReader)
	for {
		// ReadBytes与ReadString很类似
		if data, err := bufReader.ReadString(' '); err != nil {
			fmt.Println(data, err)
			break
		} else {
			fmt.Println(data, err)
		}
	}
}
