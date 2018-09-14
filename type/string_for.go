package main

import "fmt"

// Unicode VS UTF-8
// Unicode 只是一个符号集，它只规定了符号的二进制代码. 例如, 汉字严的Unicode字符是十六进制数4E25.
// 字符的二进制代码有大有小，有些可以以一个字节进行存储（比如英文字符），有些可能需要三四个字节进行存储（比如汉字）
// NOTICE: Unicode只是一个字符集（与ASCII字符集类似）， 但是Unicode并没有规定如何存储（ASCII比较好存储，就用一个字节进行存储）
// 假如每个Unicode字符都以4个字节存储（有些Unicode）, 那么传输文本中如果含有大量的英文字符时，需要传递0x00000061, 浪费网络带宽.
//
// UTF-8, UTF-16编码就是为了解决Unicode字符集的存储和传输问题。
// UTF-8 是 Unicode 的实现方式之一.
// UTF-8的特点是，有些字符是占一个字节，大多数字符占两个字节，也有字符占三或四个字节。 所以，UTF-8的特点是节省存储空间
//
// len(s) 返回的是字节数, s[i]返回了一个字节，本质上是一个byte类型(uint8).
// for range 形式返回的是一个**Unicode字符**，类型为rune(int32)，表示Unicode字符的二进制编码.
// NOTICE: rune中存储的是Unicode编码字符，并非是UTF-8编码.
// In the same way the conversion c:=[]int(s) is allowed, then each int contains a Unicode code point.
// NOTICE: []byte(s) 和 c:=[]int32(s), c:=[]int(s) 存储字符的集合

func main() {

	// NOTICE: Golang的字符串实际上是UTF-8编码的字符串
	//         在内存当中的表示形式(UTF-8编码)，61 e4 bd a0 e5 a5 bd
	//		   61 e4 bd a0 e5 a5 bd. 其中a占用一个字节，内容为61，等于ASCII字符的值；你好分别占用了三个字节.
	// s := "a你好"
	s := "\u0061\u4f60\u597d" // 直接给定UNICODE字符或者二进制形式
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x %T\n", s[i], s[i])
	}

	for i, c := range s {
		fmt.Printf("%d %x %c %T\n", i, c, c, c)
	}

	// If erroneous UTF-8 is encountered, the character is set to U+FFFD and the index advances by one byte.
	s = string([]uint8{21, 21, 255, 254, 253})
	for i, c := range s {
		fmt.Printf("%d %x %c %T\n", i, c, c, c)
	}
}
