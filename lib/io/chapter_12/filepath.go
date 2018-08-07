package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	filename := filepath.Base("/root/share")
	fmt.Println(filename)

	path, _ := filepath.Abs(".")
	fmt.Println(path)
}
