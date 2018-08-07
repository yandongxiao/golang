package main

import "os"
import "fmt"

func main() {
	var file *os.File
	var err error
	/* make sure file and err has the same scope */
	if file, err = os.Open("/tmp"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("read data")
	}
	defer file.Close()

	if files, err := file.Readdir(-1); err != nil {
		panic(err)
	} else {
		for idx := range files {
			if files[idx].IsDir() {
				fmt.Println(files[idx].Name())
			}
		}
	}

}
