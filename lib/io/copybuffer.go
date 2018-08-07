package main

import "os"
import "fmt"
import "io"

// CopyBuffer is identical to Copy except that it stages through the provided buffer (if one is required) rather than allocating a temporary one.
// If buf is nil, one is allocated; otherwise if it has zero length, CopyBuffer panics.)
func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var src *os.File
	var dest *os.File
	var err error

	src, err = os.Open("/tmp/data")
	errCheck(err)
	defer src.Close()

	dest, err = os.Create("/tmp/data.cp")
	errCheck(err)
	defer dest.Close()

	buffer := make([]byte, 1)                     // 注意，buffer只是用来缓存读的数据，最后还是要写到dest上的
	size, err := io.CopyBuffer(dest, src, buffer) // 等价于Copy
	errCheck(err)
	fmt.Println("size =", size)

}
