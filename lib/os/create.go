package main

import "os"

func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	/* no matter the data exist, return success */
	file, err := os.Create("/tmp/data")
	errcheck(err)
	defer file.Close()
}
