package main

func main() {
	str := "string"
	var pstr *string = &str

	*pstr = "------"
	println(pstr, *pstr, str)
}
