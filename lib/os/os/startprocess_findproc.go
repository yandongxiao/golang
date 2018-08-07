package main

import (
	"fmt"
	"os"
)

func main() {
	process, err := os.FindProcess(3376)
	if err != nil {
		panic(err)
	}

	err = process.Kill()
	if err != nil {
		panic(err)
	}

	fmt.Println("killed the process")
}
