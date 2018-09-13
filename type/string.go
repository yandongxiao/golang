// golang中的字符串本身就是采用UTF-8进行编码的, 而UTF-8编码中一个字符的尺寸是不确定的
// 字符串变量的内容是不能被修改的. 这个与python字符串的特性相同，可见，字符串变量应该是引用类型
// 但因为是不可变类型，所以可以像使用值的方式使用它
package main

import "fmt"

func main() {
	// 1. string的底层指针是没有办法获取的
	// 2. data 和 修改后的bd 肯定没有指向同一块内存。
	// 如果某个函数返回值是string，大可放心使用，就像使用整型一样.
	// string is immutable
	data := "hello world"
	bd := []byte(data)
	bd[0] = 'H'

	// 直接打印一个byte slice
	fmt.Printf("%q\n", bd)
	fmt.Println(bd)
}
