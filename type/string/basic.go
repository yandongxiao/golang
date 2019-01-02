package main

func main() {
	// 不能⽤用序号获取字节元素指针，&s[i] ⾮非法
	//s := "helloworld"
	//fmt.Println(&s[0])

	s2 := `
b\r\n\x00
c`
	print(s2)
}
