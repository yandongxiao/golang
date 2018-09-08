// indent 和 compact是一对相互反作用的函数
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	data := indent()
	compact(data)
}

func indent() []byte {

	var buffer bytes.Buffer
	err := json.Indent(&buffer, []byte("{\"name\":\"jack\",\"age\":10}"), "", "    ")
	checkError(err)
	fmt.Println(buffer.String())
	return buffer.Bytes()
}

// compact:紧凑之意
func compact(data []byte) {
	buffer := new(bytes.Buffer)

	// 应用一: 在发送json字符串之前，对数据进行进一步的压缩.
	err := json.Compact(buffer, data)
	checkError(err)
	fmt.Println(buffer.String())

	// 应用二: 对value进行压缩. 意义不大，可以在最后进行一并压缩
	err = json.Compact(buffer, []byte("  10 "))
	checkError(err)
	fmt.Println(buffer.String())

	err = json.Compact(buffer, []byte(" [1, 2, 3, 4, 5]  "))
	checkError(err)
	fmt.Println(buffer.String())
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
