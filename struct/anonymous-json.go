// Struct types tend to be verbose because they often involve a line for each field.
// Although we could write out the whole type each time it is needed, the repetition would get tiresome.
// Instead, struct types usually appear within the declaration of a named type like Employee
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 与 type INT int 进行类比，即可得
// struct{} 就是一个类型，Person是这个类型的别名
type Person struct {
	name string
	age  int
}

func main() {

	// Decoding deeply nested JSON data
	// NOTICE: 你再也不用为了每个HTTP的响应创建新的struct了
	f, _ := os.Open("anonymous-json.body")
	data, _ := ioutil.ReadAll(f)
	person := struct {
		Code int
		Data struct {
			Name     string
			Age      int
			Children []struct {
				Name string
				Sex  string
			}
		}
	}{}

	json.Unmarshal(data, &person)
	fmt.Println(person)
}
