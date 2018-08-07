package main

import "fmt"

// 例如："a你好"字符串的Unicode编码是"\u0061\u4f60\u597d"，对Unicode进行编码的方式有多种，如UTF-8，UTF-16等.
// UTF-8 是 Unicode 的实现方式之一
// UTF-8的特点是，有些字符是占一个字节，大多数字符占两个字节，也有字符占三或四个字节
// 所以，UTF-8的特点是节省存储空间和传输空间
// Golang的字符串实际上是UTF-8编码的字符串, 内存表示如下：
// 61 e4 bd a0 e5 a5 bd. 其中a占用一个字节，内容为61，等于ASCII字符的值；你好分别占用了三个字节，这印证了Unicode和UTF-8之间的区别

// len(s) 返回的是字节数, s[i]返回了一个字节，本质上是一个byte类型(uint8).
// for range 形式返回的是一个Unicode字符，类型为rune(int32)，表示任意一个Unicode编码的字符.
// NOTE: rune中存储的是Unicode编码字符，并非是UTF-8编码.
// If erroneous UTF-8 is encountered, the character is set to U+FFFD and the index advances by one byte.
// In the same way the conversion c:=[]int(s) is allowed, then each int contains a Unicode code point.
// []byte(s) 和 c:=[]int32(s), c:=[]int(s) 存储字符的集合

func main() {
	// s := "a你好"
	// 在内存当中的表示形式(UTF-8编码)，e4 bd a0 e5 a5 bd
	s := "\u0061\u4f60\u597d" // UNICODE编码
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x %T\n", s[i], s[i])
	}
	for i, c := range s {
		fmt.Printf("%d %x %c %T\n", i, c, c, c)
	}

	// 2 fffd � int32
	// If erroneous UTF-8 is encountered, the character is set to U+FFFD and the index advances by one byte.
	s = string([]uint8{21, 21, 255, 254, 253})
	for i, c := range s {
		fmt.Printf("%d %x %c %T\n", i, c, c, c)
	}
}

// prints: 0:ÿ 2:界
