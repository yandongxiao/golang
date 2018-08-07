package main

import "fmt"
import "io"

func myfunc(reader io.Reader) {
	rc, ok := reader.(io.ReadCloser)
	if ok {
		fmt.Println(rc)
	} else {
		fmt.Println("not ok")
	}

}

func main() {
	myfunc(nil)
}
