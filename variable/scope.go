// scope等级：
//	global level or package level
//	local level
//	structure level
//  下层可以隐藏上层变量
package main

var a = "G" // global (package) scope

func main() {
	n()
	m()
	n()
}
func n() {
	print(a)
}
func m() {
	a := "O" // new local variable a is declared
	print(a)
}

// GOG
