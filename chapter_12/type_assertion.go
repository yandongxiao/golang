package main

func main() {
	data := 10
	var val interface{}
	val = data

	if inst, ok := val.(int); ok {
		println("inst is int type", inst)
	}
}
