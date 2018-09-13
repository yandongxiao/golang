package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// Decoding deeply nested JSON data
	// NOTICE: 你再也不用为了每个HTTP的响应创建新的struct了
	f, _ := os.Open("body.json")
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
