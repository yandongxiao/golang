package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name   string
	ChildF ChildF
}

type Person2 struct {
	Name string
	ChildF
}

type Person3 struct {
	Name string
	C
}

type C struct {
	Age int
}

type ChildF struct {
	Age int
}

func ExampleEmcode() {
	p := Person{
		Name: "jack",
		ChildF: ChildF{
			Age: 10,
		},
	}
	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	// 注意：匿名 field 在 json encode的时候，没有带ChildF
	p2 := Person2{
		Name: "jack",
		ChildF: ChildF{
			Age: 10,
		},
	}
	data, _ = json.Marshal(p2)
	fmt.Println(string(data))

	// Output:
	// {"Name":"jack","ChildF":{"Age":10}}
	// {"Name":"jack","Age":10}
}

var data = `
{
    "code":200,
    "data": {
        "name": "jack",
        "age": 10,
        "children": [
            {
                "name": "alice",
                "sex": "m"
            },
            {
                "name": "bob",
                "sex": "f"
            }
        ]
    }

}
`

var data2 = `
{
    "code":200,
    "name": "jack",
    "age": 10,
    "children": [
        {
            "name": "alice",
            "sex": "m"
        },
        {
            "name": "bob",
            "sex": "f"
        }
    ]
}
`

func TestAnonymousUnmarshal(t *testing.T) {
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
	err := json.NewDecoder(strings.NewReader(data)).Decode(&person)
	assert.Nil(t, err)
	assert.True(t, person.Data.Name == "jack")
	assert.True(t, person.Data.Children[0].Name == "alice")

	// 注意：
	// 解码的时候，必须有 Data 这一层（比如 data 和 data2 的内容h），即使你有 json:",inline" 的 tag 也不行
	// 所以，这里有一个问题：匿名field在编码的时候，不会包含Data这一层；但是在解码的时候，强制需要Data这一层。
	// 解决办法：Data的类型必须在 外层 进行定义。
	type Data struct {
		Name     string
		Age      int
		Children []struct {
			Name string
			Sex  string
		}
	}
	person2 := struct {
		Code int
		Data `json:",inline"`
	}{}
	err = json.NewDecoder(strings.NewReader(data2)).Decode(&person2)
	assert.Nil(t, err)
	assert.True(t, person2.Data.Name == "jack")
	assert.True(t, person2.Data.Children[0].Name == "alice")
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
		},
		title: "The Million Pound Note",
		pages: 96,
	}
	_ = aBook
	// OUtput:
	//
}
