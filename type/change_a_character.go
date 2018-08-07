package main

func main() {
	name := "ydx"
	buff := []byte(name)
	buff[0] = 'Y'
	name = string(buff)

	println(name)
}
