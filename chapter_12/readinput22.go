package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputBuffer := bufio.NewReader(os.Stdin)
	data, err := inputBuffer.ReadString('\n') /* NOTE the difference of \n and io.EOF*/
	if err == nil {
		fmt.Print(data)
	} else if err == io.EOF {
		fmt.Println("read to io.EOF")
	} else {
		fmt.Println("read data error")
	}

}
