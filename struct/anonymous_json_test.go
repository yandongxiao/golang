package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

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

func ExampleAnonymous() {
	// Annoymous struct types are allowed used as the types of the fields of another struct type.
	// Annoymous struct type literals are also allowed to be used in composite literals.
	var aBook = struct {
		author struct { // the type of this field is an anonymous struct type
			firstName, lastName string
			gender              bool
		}
		title string
		pages int
	}{
		author: struct {
			firstName, lastName string
			gender              bool
		}{
			firstName: "Mark",
			lastName:  "Twain",
		}, // the type in the composite literal is an anonymous struct type
		title: "The Million Pound Note",
		pages: 96,
	}
	_ = aBook
	// OUtput:
	//
}
