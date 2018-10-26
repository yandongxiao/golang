//go:generate stringer -type=Pill
package main

import "fmt"

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

//go:generate echo hello
//go:generate go run painkiller.go
//go:generate  echo file=$GOFILE pkg=$GOPACKAGE
func main() {
	fmt.Println("main func")
}
