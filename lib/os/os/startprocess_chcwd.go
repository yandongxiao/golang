package main

import "os"

func main() {
	attr := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   "/tmp",
	}
	var args []string = []string{"/tmp/show_cwd.sh"}
	os.StartProcess(args[0], args, &attr)
}
