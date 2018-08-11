package main

import (
	"fmt"
	"os"
)

func main() {

	// If there is an error, it will be of type *PathError.
	err := os.Chdir("/tmpd")
	if err != nil {
		pathErr := err.(*os.PathError)
		fmt.Println(pathErr.Op)   // chdir
		fmt.Println(pathErr.Path) // /tmpd
		fmt.Println(pathErr.Err)  // no such file or directory
		fmt.Println(pathErr)      // chdir /tmpd: no such file or directory
	}
}
