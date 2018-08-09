## Scanner

### Scanner也缓存数据

```
Scanner provides a convenient interface for reading data such as a file
of newline-delimited lines of text. 所以，Scanner也是通过缓存数据，提供方便的接口。

type Scanner struct {
	// 缓存相关
	r            io.Reader // The reader provided by the client.
	buf          []byte    // Buffer used as argument to split.
	start        int       // First non-processed byte in buf.
	end          int       // End of data in buf.
	err          error     // Sticky error.
	
	split        SplitFunc // The function to split the tokens.
	maxTokenSize int       // Maximum size of a token; modified by tests.
	token        []byte    // Last token returned by split.
	
	// Scan panics if the split function returns 100 empty tokens without
	// advancing the input. This is a common error mode for scanners.
	empties      int       // Count of successive empty tokens.
	// 一旦调用Scan方法，scanCalled置为true.
	scanCalled   bool      // Scan has been called; buffer is in use.
	done         bool      // Scan has finished.
}
```

### 何谓token？

```
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
```

SplitFunc is the signature of the split function used to tokenize the input. 
将底层r返回的数据流类比于TCP返回的数据流，上层应用需要的是一个个可拆分的独立的消息。SplitFunc的工作就是数据流消息化。

输入：

The arguments are an initial substring of the remaining unprocessed data 
and a flag, atEOF, that reports whether the Reader has no more data to give.

输出：

The return values are the number of bytes to advance the input and the next token to return to the user, plus an error

例子：

If the data does not yet hold a complete token, for instance if it has no 
newline while scanning lines, SplitFunc can return (0, nil, nil) to 
signal the Scanner to read more data into the slice and try again with 
a longer slice starting at the same point in the input.

Successive calls to the Scan method will step through the 'tokens' of a file,
skipping the bytes between the tokens. 夹杂在两个token之间的数据会被丢弃。

The specification of a token is defined by a split function of
type SplitFunc; the default split function breaks the input into lines
with line termination stripped.  默认是读一行

Split functions are defined in this package for scanning a file into lines, bytes, UTF-8-encoded runes, and space-delimited words. The client may instead provide a custom split function.

### Scanner如何工作？

一般是对Scan方法进行遍历

```
func (s *Scanner) Scan() bool
```

Scan advances the Scanner to the next token, which will then be available through the Bytes or Text method. 

Scanning stops unrecoverably at EOF, the first I/O error, or a token too
large to fit in the buffer. When a scan stops, the reader may have
advanced arbitrarily far past the last token. 此时Scan方法返回false

Scan panics if the split function returns 100 empty tokens without advancing the input. This is a common error mode for scanners. 所以，Scanner还是比较任性的，调用者对它的错误控制能力较差。

Programs that need more control over error handling or large tokens, or must run sequential scans on a reader, should use bufio.Reader instead.