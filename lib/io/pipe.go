package main

import (
	"fmt"
	"io"
	"time"
)

func main() {
	// Pipe creates a synchronous in-memory pipe.
	// 何谓同步?
	// Read: it reads data from the pipe, blocking until a writer arrives or the write end is closed.
	// Write:it writes data to  the pipe, blocking until readers have consumed all the data or the read end is closed.
	// there is no internal buffering
	// 从上面的描述，感觉pipe像是chan实现的.
	//
	// 关于并发：
	// It is safe to call Read and Write in parallel with each other or with Close.
	// Parallel calls to Read, and parallel calls to Write, are also safe: the individual calls will be gated sequentially.
	//
	// 一次读操作是不会跨越两个写操作之间的数据的
	// Reads on one end are matched with writes on the other, copying data directly between the two;
	//
	// 关于Close：
	//	Close closes the reader; subsequent writes to the write half of the pipe will return the error ErrClosedPipe.
	//	Close closes the writer; subsequent reads from the read half of the pipe will return no bytes and EOF.
	//  Close will complete once pending I/O is done(尤其是写)
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "hello")
		fmt.Println("write hello done")
		fmt.Fprint(w, "world\n")
		fmt.Println("write world done")
		w.Close()
	}()

	data := make([]byte, 3)
	time.Sleep(time.Second)
	fmt.Println(r.Read(data))
	fmt.Println(r.Read(data))
	fmt.Print(r.Read(data))
	fmt.Print(r.Read(data))
}
