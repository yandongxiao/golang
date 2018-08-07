package main

import "os"

func main() {
	attr := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	var args []string = []string{"/tmp/helloworld.sh"}
	os.StartProcess(args[0], args, &attr)
}
