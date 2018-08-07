// zero-value 值类型的默认值为0，false或空字符串等，引用类型的默认值为nil
//
// 目前slice，map、chan是引用类型，证明指针和函数也属于引用类型
package main

import "fmt"

func main() {
	var myfunc func(int) int
	if myfunc == nil {
		fmt.Println("函数属于引用类型")
	}

	var p *int
	if p == nil {
		fmt.Println("指针属于引用类型")
	}

	var i interface{}
	if i == nil {
		fmt.Println("接口属于引用类型")
	}
}
