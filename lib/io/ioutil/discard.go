package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// var Discard io.Writer = devNull(0))
// Discard is an io.Writer on which all Write calls succeed without doing anything.
func main() {
	fmt.Println(io.Copy(ioutil.Discard, strings.NewReader("helloworld")))
}
