// rune is an alias for int32 and is equivalent to int32 in all ways.
// It is used, by convention, to distinguish character values from integer values.
// type rune = int32 这种语法**只能**出现在builtin.go中.
// type byte = uint8 只此两种
package main

func test1() {}
func test2() {}

func main() {
	// NOTE: 不能将rune等价为type rune int32
	println("------")
	m := int32(1)
	n := rune(1)
	println(m == n) //true
	f = m
	g = n
	println(f == g) //true

	type INT32 int32
	q := INT32(1)
	//println(m == q) //panic
	g = q
	println(f == g) // false
}
