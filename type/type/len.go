// NOTE: 对于自定义的类型，可以调用len，append等函数
// func len(v Type) int
// func append(slice []Type, elems ...Type) []Type
// type Type int
//    Type is here for the purposes of documentation only. It is a stand-in
//    for **any Go type**, but represents the same type for any given function invocation.
package main

type Stack []interface{}

func main() {
	var st1 Stack
	println(len(st1))
	st1 = append(st1, 100)
	println(len(st1))
}
