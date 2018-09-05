package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	r, w, err := os.Pipe()
	rw, _ := os.Create("/tmp/output")
	if err != nil {
		fmt.Fprintf(rw, "pipe: %s\n", err)
		return
	}

	p, err := os.StartProcess("/bin/date", []string{"date"}, &os.ProcAttr{Files: []*os.File{nil, w, w}})
	defer r.Close()
	w.Close()
	if err != nil {
		fmt.Fprintf(rw, "fork/exec: %s\n", err)
		return
	}

	io.Copy(rw, r)
	wait, err := p.Wait()
	if err != nil {
		fmt.Fprintf(rw, "wait: %s\n", err)
		return
	}

	if !wait.Exited() {
		fmt.Fprintf(rw, "date: %v\n", wait)
		return
	}
}
