# io

Package io provides **basic** interfaces to **I/O primitives**. package io做了以下工作：

1. 抽象出IO原语，比如Read操作的原语```Read(p []byte) (n int, err error)```；
2. 对IO原语进行接口封装，使得IO操作可以有**多态**的行为；
3. 基于IO接口提供**非常通用**的IO操作。

## 一流的公司做标准

Package io provides **basic** interfaces to **I/O primitives**. 几乎所有的系统都存在IO操作，比如文件系统，网络系统等，io package**高度抽象**出了这些系统共有的特点，形成了IO primitives。

每一个IO primitive又被封装成一个interface。

一流的公司做标准，io package就是这样一流的package。我们io.Reader为例

```
type Reader interface {
    Read(p []byte) (n int, err error)
}

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
considering the error err. Doing so correctly handles I/O errors that
happen after reading some bytes and also both of the allowed EOF
behaviors.

Implementations of Read are discouraged from returning a zero byte count
with a nil error, except when len(p) == 0. Callers should treat a return
of 0 and nil as indicating that nothing happened; in particular it does
not indicate EOF.

Implementations must not retain p.
```



>
> 具体怎么实现Read接口，io.Reader并不关心，但是你的实现，必须满足Read操作规范！

## error

package io 不但封装了IO原语，同时为实现者提供了**标准的**错误标识。比如，Read操作肯定会读取到末尾，那么所有的实现者都应该返回io.EOF，而非自定义。

毕竟，package io中只会定义**最通用**的错误标识，每个实现者可以根据需求自定义一些错误。如：bufio

```
ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
ErrBufferFull        = errors.New("bufio: buffer full")
ErrNegativeCount     = errors.New("bufio: negative count")
```
>
> 注意，错误并非是常量，因为常量只能应用于primitive type(数字、布尔、字符串)


## functions


## pipe

# 读
type Reader interface {
    func Read(p []byte) (n int, err error)
}
Reader is the interface that wraps the basic Read method

Read reads up to len(p) bytes into p
It returns the number of bytes read (0<=n<=len(p)) and any error encounterd.

# 不能假设[n, len(p))之间的数据没有被更改
Even if Read returns n < len(p), it may use all of p as scratch space(临时空间) during the call.

# 即使正确返回的情况下，并不一定返回len(p)个字节. 但是这说明不了Read是阻塞还是非阻塞式
If some data is available but not len(p) bytes, Read conventionally returns what is available instead of waiting for more.

# 遇到错误的情况下，错误也有可能是在下一次调用后返回
When Read encounters an error or end-of-file condition after successfully reading n > 0 bytes, it returns the number of bytes read.
It may return the non-nil error from the same call or return the error (and n==0) from a subsequent call.
An instance of this general case is that a Reader returning a non-zero number of bytes at the end of the input stream may return 
err == EOR or err == nil. The next Read should return 0, EOF.

# 失败的情况下也要检查n的值，n>0就表示正确读取了这些数据
Caller should always process the n > 0 bytes returned brfore considering the error err.

Implementations of Read are discouraged from returning n=0 and err=nil except when len(p)=0
Caller should treat a return of 0 and nil as indicating that nothing happened. NOTE: It does not indicate EOF


# 写
type Writer interface {
    func Write(p []byte) (n int, err error)
}

Writer is the interface that wrpas the basic Write method

Write writes len(p) bytes from p to the underlying data stream. It returns the number of bytes written from p (0<=n<=len(p)) and any error encountered
that cased the write to stop early.
Write must return a non-nil error if it returns n < len(p)
Write must not modify the slice data

