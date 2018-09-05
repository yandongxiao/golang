package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data, _ := json.Marshal("nihao")
	fmt.Println(string(data))
}
