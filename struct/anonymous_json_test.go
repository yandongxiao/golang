package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 与 type INT int 进行类比
// type JPerson struct{...}
// struct{} 本身就是一个类型
// struct{}{} 不但定义该类型，同时对它进行初始化
type JPerson struct {
	name string
	age  int
}

func ExampleAnonymousUnmarshalJson() {
	// Decoding deeply nested JSON data
	// NOTICE: 你再也不用为了每个HTTP的响应创建新的struct了
	f, _ := os.Open("anonymous_json.body")
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
	// Output:
	// {200 {jack 10 [{alice m} {bob f}]}}
}
