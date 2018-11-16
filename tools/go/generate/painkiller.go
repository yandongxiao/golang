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

// NOTE: 不能将所有的generate指令解释为一个大的脚本！
// 如下，在第一条指令中设置hello, 但是在第二条指令中并没有获取到它的值
//go:generate -command ECHO echo nihao
//go:generate EHCO world
//go:generate echo $hello, $PATH
//go:generate go run painkiller.go
//go:generate  echo file=$GOFILE pkg=$GOPACKAGE
func main() {
	a := 1
	fmt.Println("main func", a)
}
