// 字符串变量的内容是不能被修改的. 所以可以像使用值的方式使用它
package main

import "fmt"

func main() {
	// 1. string的底层指针是没有办法获取的
	// 2. data 和 修改后的bd 肯定没有指向同一块内存。
	// 3. 以下也是修改string的方式.
	data := "hello world"
	bd := []byte(data)
	bd[0] = 'H'

	// 直接打印一个byte slice
	fmt.Printf("%q\n", bd)
	fmt.Println(data)
}
