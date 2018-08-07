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
		fmt.Println("open directory succeed")
	}
	defer file.Close()

	if err = file.Chdir(); err != nil {
		fmt.Println(err)
		return
	}

	if file, err = os.Open("data2"); err != nil { /* not a abs path */
		fmt.Println(err)
	} else {
		fmt.Println("use a relative path to open")
	}
	defer file.Close()
}
