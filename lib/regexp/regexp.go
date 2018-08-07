package main

import "regexp"

func main() {
	match, _ := regexp.MatchString("llo", "helloworld")
	println(match)
}
