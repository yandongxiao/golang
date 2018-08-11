# pipe

```io.pipe```与网络编程中的管道编程是相同的。共同特点如下：

1. 创建一个pipe，返回管道的两端，一端负责读，一端负责写；
2. 一个线程负责写，另一个线程负责读；
3. 数据是流式传输，需要上层应用识别为消息。

## 实现原理

[linux管道pipe详解](https://blog.csdn.net/oguro/article/details/53841949)介绍了Linux操作系统下pipe的工作原理。 总结如下：

```
读管道：
	1. 管道中有数据，read返回实际读到的字节数。
	2. 管道中无数据：
		(1) 管道写端被全部关闭，read返回0 (好像读到文件结尾)
		(2) 写端没有全部被关闭，read阻塞等待(不久的将来可能有数据递达，此时会让出cpu)
写管道：
	1. 管道读端全部被关闭， 进程异常终止(也可使用捕捉SIGPIPE信号，使进程不终止)
	2. 管道读端没有全部关闭：
		(1) 管道已满，write阻塞。
		(2) 管道未满，write将数据写入，并返回实际写入的字节数。
```

golang的pipe的数据结构如下:

```
type pipe struct {
	rl    sync.Mutex // gates readers one at a time. 为所有读操作排序
	wl    sync.Mutex // gates writers one at a time. 为所有写操作排序
	// 互斥量l + 条件变量rwait，wwait，完成读写操作的协同工作
	l     sync.Mutex // protects remaining fields
	rwait sync.Cond  // waiting reader
	wwait sync.Cond  // waiting writer
	data  []byte     // data remaining in pending write 读写的数据来源
	rerr  error      // if reader closed, error to give writes
	werr  error      // if writer closed, error to give reads
}

```

golang pipe 划重点：

1. Pipe creates a **synchronous** in-memory pipe. 所以linux pipe像是有缓存的管道，而golang pipe是无缓存的管道（读写两端必须同时在线）。
2. ```data```只是读写数据的来源，不是用来缓存读写数据的。它与```func (w *PipeWriter) Write(data []byte) (n int, err error)```Write操作的输入参数```data```，指向了同一块内存。
3. Close will complete once pending I/O is done. 所以```close```操作也不是说立即生效，必须在当前读写操作完成之后。
4. Parallel calls to Read, and parallel calls to Write are safe。Linux下的文件描述符可不是线程安全的。

>
> 总结下来，线程安全，同步，write和read操作是一对多的关系。
> Linux下write和read操作是多对多的关系，也可以说是没关系。
