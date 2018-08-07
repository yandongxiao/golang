package main

import "os"

func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	attr := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	var args []string = []string{"/bin/sleep", "10000"}

	process, err := os.StartProcess(args[0], args, &attr)
	errcheck(err)

	err = process.Signal(os.Kill)
	errcheck(err)
}
