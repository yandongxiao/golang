package main

import "log"
import "os"

func main() {
	log.Output(12, "output data")

	file, _ := os.Create("/tmp/mylog")
	log.SetOutput(file)
	log.Output(12, "output data")
}
