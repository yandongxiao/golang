package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	procAttr := &os.ProcAttr{
		Dir: "/tmp",
	}

	argv := []string{"/bin/sleep", "1"}
	proc, err := os.StartProcess(argv[0], argv, procAttr)
	checkError(err)

	// The only signal values guaranteed to be present on all systems are Interrupt and kill
	//proc.Signal(os.Kill)
	//proc.Signal(os.Interrupt)
	proc.Kill()

	stat, err := proc.Wait()
	checkError(err)
	fmt.Println(stat)
}
