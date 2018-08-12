# strings

## as a reader

```
// A Reader implements the io.Reader, io.ReaderAt, io.Seeker, io.WriterTo,
// io.ByteScanner, and io.RuneScanner interfaces by reading
// from a string.
type Reader struct {
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}

// NewReader returns a new Reader reading from s.
// It is similar to bytes.NewBufferString but more efficient and read-only.
func NewReader(s string) *Reader
```

## as a replacer


```
// NewReplacer returns a new Replacer from a list of old, new string pairs.
// Replacements are performed in order, without overlapping matches.
func NewReplacer(oldnew ...string) *Replacer

It is safe for concurrent use by multiple goroutines.
```

给定一堆的(old, new)字符串列表，用于替换。

## 通用方法

注意与bytes的方法进行类比，可谓是**一一对应**。

```
func Compare(a, b string) int
func Contains(s, substr string) bool
func ContainsAny(s, chars string) bool
func ContainsRune(s string, r rune) bool
func Count(s, sep string) int

func EqualFold(s, t string) bool

func Fields(s string) []string
func FieldsFunc(s string, f func(rune) bool) []string

func HasPrefix(s, prefix string) bool
func HasSuffix(s, suffix string) bool

func Index(s, sep string) int
func IndexAny(s, chars string) int
func IndexByte(s string, c byte) int
func IndexFunc(s string, f func(rune) bool) int
func IndexRune(s string, r rune) int

func Join(a []string, sep string) string

func LastIndex(s, sep string) int
func LastIndexAny(s, chars string) int
func LastIndexByte(s string, c byte) int
func LastIndexFunc(s string, f func(rune) bool) int

func Map(mapping func(rune) rune, s string) string

func Repeat(s string, count int) string
func Replace(s, old, new string, n int) string

func Split(s, sep string) []string
func SplitAfter(s, sep string) []string
func SplitAfterN(s, sep string, n int) []string
func SplitN(s, sep string, n int) []string

func Title(s string) string
func ToLower(s string) string
func ToLowerSpecial(_case unicode.SpecialCase, s string) string
func ToTitle(s string) string
func ToTitleSpecial(_case unicode.SpecialCase, s string) string
func ToUpper(s string) string
func ToUpperSpecial(_case unicode.SpecialCase, s string) string

func Trim(s string, cutset string) string
func TrimFunc(s string, f func(rune) bool) string
func TrimLeft(s string, cutset string) string
func TrimLeftFunc(s string, f func(rune) bool) string
func TrimPrefix(s, prefix string) string
func TrimRight(s string, cutset string) string
func TrimRightFunc(s string, f func(rune) bool) string
func TrimSpace(s string) string
func TrimSuffix(s, suffix string) string
```


