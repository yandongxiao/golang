package main

import "fmt"
import "io/ioutil"

// var Discard io.Writer = devNull(0))
// Discard is an io.Writer on which all Write calls succeed without doing anything.
func main() {
	fmt.Println(ioutil.Discard.Write([]byte("helloworld")))
}
