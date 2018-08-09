# bytes

该package可分为三部分：Functions, Reader, Buffer.

## Functions

关键字：Compare, Contain, Equal, Fields, Has, Index, Map, Replace, Split, Trim, To

```
func Compare(a, b []byte) int
func Contains(b, subslice []byte) bool
func ContainsAny(b []byte, chars string) bool
func ContainsRune(b []byte, r rune) bool
func Count(s, sep []byte) int
func Equal(a, b []byte) bool
func EqualFold(s, t []byte) bool
func Fields(s []byte) [][]byte
func FieldsFunc(s []byte, f func(rune) bool) [][]byte
func HasPrefix(s, prefix []byte) bool
func HasSuffix(s, suffix []byte) bool
func Index(s, sep []byte) int
func IndexAny(s []byte, chars string) int
func IndexByte(s []byte, c byte) int
func IndexFunc(s []byte, f func(r rune) bool) int
func IndexRune(s []byte, r rune) int
func Join(s [][]byte, sep []byte) []byte
func LastIndex(s, sep []byte) int
func LastIndexAny(s []byte, chars string) int
func LastIndexByte(s []byte, c byte) int
func LastIndexFunc(s []byte, f func(r rune) bool) int
func Map(mapping func(r rune) rune, s []byte) []byte
func Repeat(b []byte, count int) []byte
func Replace(s, old, new []byte, n int) []byte
func Runes(s []byte) []rune
func Split(s, sep []byte) [][]byte
func SplitAfter(s, sep []byte) [][]byte
func SplitAfterN(s, sep []byte, n int) [][]byte
func SplitN(s, sep []byte, n int) [][]byte
func Title(s []byte) []byte
func ToLower(s []byte) []byte
func ToLowerSpecial(_case unicode.SpecialCase, s []byte) []byte
func ToTitle(s []byte) []byte
func ToTitleSpecial(_case unicode.SpecialCase, s []byte) []byte
func ToUpper(s []byte) []byte
func ToUpperSpecial(_case unicode.SpecialCase, s []byte) []byte
func Trim(s []byte, cutset string) []byte
func TrimFunc(s []byte, f func(r rune) bool) []byte
func TrimLeft(s []byte, cutset string) []byte
func TrimLeftFunc(s []byte, f func(r rune) bool) []byte
func TrimPrefix(s, prefix []byte) []byte
func TrimRight(s []byte, cutset string) []byte
func TrimRightFunc(s []byte, f func(r rune) bool) []byte
func TrimSpace(s []byte) []byte
func TrimSuffix(s, suffix []byte) []byte
```

## Reader

```
func NewReader(b []byte) *Reader
    NewReader returns a new Reader reading from b.
```
以读的方式对外暴露byte slice的访问方式，由于它并不关联设备，且只以``b```的数据作为“设备”的全部内容
所以bytes.Reader的内容可以认为已经全部load到了内存，所以bytes.Reader可以提供更丰富的操作。

```
A Reader implements the io.Reader, io.ReaderAt, io.WriterTo, io.Seeker,
io.ByteScanner, and io.RuneScanner interfaces by reading from a byte
slice. Unlike a Buffer, a Reader is read-only and supports seeking.
```

## Buffer

与Reader相比，Buffer不仅提供了读操作接口，也支持写操作。Buffer的一个应用是聚合多个string，Buffer.WriteString.

```
// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
// The zero value for Buffer is an empty buffer ready to use.
type Buffer struct {
	buf       []byte            // contents are the bytes buf[off : len(buf)]
	off       int               // read at &buf[off], write at &buf[len(buf)]
	runeBytes [utf8.UTFMax]byte // avoid allocation of slice on each call to WriteRune
	bootstrap [64]byte          // memory to hold first slice; helps small buffers avoid allocation.
	lastRead  readOp            // last read operation, so that Unread* can work correctly.
}
```

```buf```用于缓存数据，[0, off）表示已经读取的数据，[off, len(buf)]表示未读取的数据；[len(buf), cap(buf)]表示要写的数据的位置。
Buffer还可以自动扩展buf，以适应写入了大量的数据。
