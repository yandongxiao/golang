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

- 定义的标准列表

	1. ```Read(p []byte) (n int, err error)```
	2. 	```Write(p []byte) (n int, err error)```
	3. ```Close() error```
	4. ```	ReadAt(p []byte, off int64) (n int, err error)```
	5. ```	ReadFrom(r Reader) (n int64, err error)```
	6. ```WriteString(s string) (n int, err error)```
	7. ```	WriteAt(p []byte, off int64) (n int, err error)```
	8. ```	WriteTo(w Writer) (n int64, err error)```
	9. ```	Seek(offset int64, whence int) (int64, error)```

>
> 以上所有方法都有对应的接口，如Read操作的封装接口就是Reader.

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

## 二次封装

基于IO原语，golang也提供了二次封装的应用类型。包括：

### interface

```
type LimitedReader struct {
	R Reader // underlying reader
	N int64  // max bytes remaining
}
```

```
type SectionReader struct {
	r     ReaderAt
	base  int64
	off   int64
	limit int64
}

// NewSectionReader returns a SectionReader that reads from r
// starting at offset off and stops with EOF after n bytes.
func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader {
	return &SectionReader{r, off, off, off + n}
}

// TeeReader returns a Reader that writes to w what it reads from r.
// All reads from r performed through it are matched with
// corresponding writes to w. There is no internal buffering -
// the write must complete before the read completes.
// Any error encountered while writing is reported as a read error.
type teeReader struct {
	r Reader
	w Writer
}
```

### functions

```
func Copy(dst Writer, src Reader) (written int64, err error)
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
func ReadFull(r Reader, buf []byte) (n int, err error)
func WriteString(w Writer, s string) (n int, err error)
```

### functions from ioutil

ioutil的IO操作大多是与文件系统有关，即用户无需自己打开，读，关闭等操作，一键搞定

```
创建文件：
func TempDir(dir, prefix string) (name string, err error)
    Multiple programs calling TempDir simultaneously will not choose the same directory.
func TempFile(dir, prefix string) (f *os.File, err error)

读写文件:
func ReadDir(dirname string) ([]os.FileInfo, error)
func ReadFile(filename string) ([]byte, error)
func WriteFile(filename string, data []byte, perm os.FileMode) error
var Discard io.Writer = devNull(0)

func NopCloser(r io.Reader) io.ReadCloser
func ReadAll(r io.Reader) ([]byte, error)
    注意与func ReadFull(r Reader, buf []byte) (n int, err error)的区别
```
