package main

import "strings"
import "fmt"

func main() {
	replcaer := strings.NewReplacer("a", "A", "b", "B")

	fmt.Println(replcaer.Replace("aabbnia"))
}
