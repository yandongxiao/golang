## Reader

[IO常识](https://www.ibm.com/developerworks/cn/linux/l-cn-directio/index.html)

关键字如下：``` synchronous writes```, ```deferred writes```, ``` asynchronous writes```, ```self-caching applications```, ```内存映射```, ```直接IO, O_DIRECT```, ```异步访问文件```。

bufio package 实现了io.Reader和io.Writer两个接口，为IO操作提供缓存功能，提高读写性能。

> 注意不一定会明显提高性能，因为Linux的默认IO操作都会内核的页缓存。

类型定义：

```
bufio package
type Reader struct {
	buf 	[]byte
	rd 		io.Reader
	r, w 	int
	err		error
	lastByte int
	lastRuneSize int
}
```

### Read操作

```
io package
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

从Reader的数据结构上分析，bufio reader 一次从底层```rd```读取多个字节（默认值是4096），根据Read方法的输入参数p，调用n = copy(p, buf)并返回。

bufio的缓存```buf```中存储了调用者尚未读取的部分[r, w)，和最近已经读取的部分[0, r)。

如果底层```rd```读取失败，则err记录了失败的原因。

> 注意
> rd.Read()操作即使调用失败（err != nil），n值也是有意义的。即：1. 不能置为负数；2. 如果n>0，表示本次调用读取到的正确的数据量。

- 从底层```rd```读取数据的时机

	1. ```r == w ``` 表示缓存为空
	2. ```err == nil```	保存底层rd上次的Read操作的error的返回值

	考虑以下场景：底层rd上次读取操作返回n=100, err=EOF, 但是调用者len(p) == 10，所以站在bufio的角度，本次Read操作应该返回```(10, nil)```。如果接下来的Read调用, len(p)大于90，那么Read操作也只是返回```(90, nil)```, 毕竟这样的返回也是符合io.Reader的语义规范。第三次调用Read，此时就需要将**err抛给用户**，让他来处理！
	
	即：**底层rd的Read操作的任何error都会向上抛给调用者**。
	
- 如何处理```len(p) > len(buf)```

	无论bufio.Reader如何处理，都必须满足io.Reader的规范。规范如下：

	```
	 Read reads up to len(p) bytes into p. It returns the number of bytes
    read (0 <= n <= len(p)) and any error encountered. Even if Read returns
    n < len(p), it may use all of p as scratch space during the call. If
    some data is available but not len(p) bytes, Read conventionally returns
    what is available instead of waiting for more.
    
	 When Read encounters an error or end-of-file condition after
    successfully reading n > 0 bytes, it returns the number of bytes read.
    It may return the (non-nil) error from the same call or return the error
    (and n == 0) from a subsequent call. An instance of this general case is
    that a Reader returning a non-zero number of bytes at the end of the
    input stream may return either err == EOF or err == nil. The next Read
    should return 0, EOF.
    
    Callers should always process the n > 0 bytes returned before
    considering the error err.
    
    Implementations of Read are discouraged from returning a zero byte count
    with a nil error, except when len(p) == 0. Callers should treat a return
    of 0 and nil as indicating that nothing happened; in particular it does
    not indicate EOF.
	```
	所以，如果缓存中有数据，直接返回缓存中的数据。如果缓存中没有数据，则直接透传调用底层rd的Read方法。

- 如何缓存数据

	1. 已缓存数据统一向左移，copy(b.buf, b.buf[b.r:b.w])；
	2. 底层rd如果返回error，则使用bufio.Reader.err记录本次错误。

### 附加操作

- Discard

	```
	func (b *Reader) Discard(n int) (discarded int, err error)
	
	Discard skips the next n bytes, returning the number of bytes discarded. 
	If Discard skips fewer than n bytes, it also returns an error. 
	If 0 <= n <= b.Buffered(), Discard is guaranteed to succeed without 
	reading from the underlying io.Reader.
	```
	所以，Discard会“尽可能”的丢弃掉指定的n个字节，如果失败（来自底层rd的Read操作），也会将失败原因抛出来。即使err!=nil, discarded的值也是有意义的。

- Peek

	```
	 func (b *Reader) Peek(n int) ([]byte, error)
    
    Peek returns the next n bytes without advancing the reader. The bytes
    stop being valid at the next read call. If Peek returns fewer than n
    bytes, it also returns an error explaining why the read is short. The
    error is ErrBufferFull if n is larger than b's buffer size.
	```
	Peek操作的特点限制了它只能预读取len(buf)个字节，否则bufio.Reader.buf无法缓存rd读取上来的数据。

- ReadByte

	```
	func (b *Reader) ReadByte() (byte, error)
	```

- ReadSlice

	```
	func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
	
	ReadSlice reads until the first occurrence of delim in the input,
   returning a slice pointing at the bytes in the buffer. The bytes stop
   being valid at the next read. If ReadSlice encounters an error before
   finding a delimiter, it returns all the data in the buffer and the error
   itself (often io.EOF). ReadSlice fails with error ErrBufferFull if the
    buffer fills without a delim. ReadSlice returns err != nil if and
    only if line does not end in delim.
	```
	它将bufio.Reader.buf的一个子slice给返回回去了，所以，下一次Read操作会覆盖返回的slice值。注意，即使读取失败，byte的值也是有意义的。

- ReadBytes / ReadString

	与ReadSlice相比，数据是独立拷贝了一份新的，所以该操作是“尽可能”地读取数据（调用ReadSlice）。如果读取到文件末尾，仍没有找到delim，io.EOF作为error的值，返回。
	
- ReadLine

	```
	func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
	```
	不建议使用，因为line是bufio.Reader.buf的一个子slice，如果届时没有读取到"\n", 它会将isPrefix置为false，并返回。
