package main

import "os"

//import "fmt"

/* *int can not be converted to uintptr */
/*
func test() {
	num := 10
	p := uintptr(&num)
	fmt.Println(p)
}
*/

func main() {
	file, _ := os.Create("/tmp/data")
	defer file.Close()

	/* NOTE: go do not create a file named mydata */
	newfile := os.NewFile(file.Fd(), "mydata")
	defer newfile.Close()
}
