package main

import (
	"container/list"
	"fmt"
)

func main() {
	scores := list.New()
	scores.PushBack(100)
	scores.PushBack(200)
	scores.PushBack(300)
	for elm := scores.Front(); elm != nil; elm = elm.Next() {
		fmt.Println(elm.Value)
	}
}
