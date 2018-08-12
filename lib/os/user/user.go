package main

import (
	"fmt"
	"os/user"
)

func main() {
	user, _ := user.Current()

	fmt.Println(user.Name)
	fmt.Println(user.Uid)
	fmt.Println(user.Gid)
	fmt.Println(user.Username)
	fmt.Println(user.HomeDir)
	fmt.Println(user.GroupIds())
}
