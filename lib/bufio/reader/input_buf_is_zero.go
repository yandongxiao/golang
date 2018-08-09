package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	buf := []byte{}
	fmt.Println(len(buf))
	fmt.Println(reader.Read(buf)) // 0 <nil>
}
