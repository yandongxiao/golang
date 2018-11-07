package main

import (
	"fmt"
	"reflect"
)

type User struct {
	UserId   int    `用户ID`
	UserName string `用户名称`
}

func main() {
	jack := User{
		UserId:   123,
		UserName: "jack",
	}

	tv := reflect.ValueOf(jack).Type()
	for i := 0; i < tv.NumField(); i++ {
		tf := tv.Field(i)
		fmt.Println(tf.Tag)
	}
}
