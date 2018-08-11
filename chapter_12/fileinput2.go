package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	input, err := os.Open("input.dat")
	if err != nil {
		fmt.Println("Open failed")
		return
	}
	defer input.Close()

	inputReader := bufio.NewReader(input)
	for {
		data, err := inputReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("read to the end of file")
			} else {
				fmt.Println("read error")
			}
			break
		}
		fmt.Print(data)
	}
}
