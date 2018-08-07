package main

import "log"

func main() {
	log.SetPrefix("prefix")
	log.Output(12, "output data")
	log.Fatalln("output data")
	log.Output(12, "output data")
}
