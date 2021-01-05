// indent 和 compact是一对相互反作用的函数
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ExampleIndentAndCompact() {
	var buffer bytes.Buffer
	err := json.Indent(&buffer, []byte("{\"name\":\"jack\",\"age\":10}"), "", "    ")
	checkError(err)
	fmt.Println(buffer.String())
}

// compact:紧凑之意
func ExampleCompact() {
	buffer := new(bytes.Buffer)
	// 对value进行压缩. 意义不大，可以在最后进行一并压缩
	err := json.Compact(buffer, []byte("  10 "))
	checkError(err)
	fmt.Println(buffer.String())

	err = json.Compact(buffer, []byte(" [1, 2, 3, 4, 5]  "))
	checkError(err)
	fmt.Println(buffer.String())
	// Output:
	// 10
	// 10[1,2,3,4,5]
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
