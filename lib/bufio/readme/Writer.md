## Writer

### Write接口定义

```
io.Reader
Write(p []byte) (n int, err error)

Write writes len(p) bytes from p to the underlying data stream. It
returns the number of bytes written from p (0 <= n <= len(p)) and any
error encountered that caused the write to stop early. Write must return
a non-nil error if it returns n < len(p). Write must not modify the
slice data, even temporarily.
```

从接口的定义上看，Reader与Writer的区别在于：Reader读取到的字节数少于调用者期望的字节数，完全没问题，err=nil；Writer写操作则必须将```p```中的所有数据写进去，否则认为是写失败，err!=nil.

### bufio.Write操作

```
type Writer struct {
	wr 		io.Writer
	buf 	[]byte
	n 		int				// 已缓存的持久化的数据量
	err 	error			// 写错误
}
```

从数据结构上看，bufio.Writer首先会比较待写数据量和```buf```的大小，如果buf比较小且n==0，则直接调用底层的```wr```的Write方法。**bufio确保一次写操作的大小至少为len(buf)，默认值为4096**。

- 如何处理```n + len(p) > len(buf)```

	bufio.Writer首先从p中取出一部分数据，填补buf缓冲区，并调用flush进行写磁盘操作；p剩余部分由于不能填充一个完整的buf缓冲区，就暂时拷贝至buf, 并返回。

- 如何处理error

	```
	If an error occurs writing to a Writer, no more data will be accepted
	and all subsequent writes will return the error.
	```
	Read操作返回错误以后，相关方法会将bufio.Reader.err重置为nil；而且Write操作并没有将bufio.Writer.err置为nil的接口！！ 

>
> 注意，只是bufio.Writer.Write的个体行为, io.Writer.Write没有这个要求

- 如何处理flush

	```
	func (b *Writer) Flush() error
	```
	调用底层wr的Write方法，持久化数据，数据来源于bufio.Writer.buf。如果写失败，bufio不但返回了失败的情况，同时会更新bufio.Writer.n和bufio.Writer.err。


### ReadFrom


```
type ReaderFrom interface {
    ReadFrom(r Reader) (n int64, err error)
}
ReadFrom reads data from r until EOF or error. The return value n is the
number of bytes read. Any error except io.EOF encountered during the
read is also returned.
注意对io.EOF的处理
```

处理逻辑仍然是：1. 读取数据到bufio.Writer.buf；2. flush数据到磁盘；3. 重复1和2，直到对象r发生读操作错误。
