package main

import "os"
import "fmt"

func main() {

	fmt.Println(os.Getenv("PATH"))

	os.Setenv("NAME", "JH")

	fmt.Println(os.Environ())

	fmt.Println(os.Clearenv())
	fmt.Println(os.Environ())
}
