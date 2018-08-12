package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	replacer := strings.NewReplacer("hello", "world")
	fmt.Println(replacer.Replace("helloworld"))
	replacer.WriteString(os.Stdout, "helloworld")
}
