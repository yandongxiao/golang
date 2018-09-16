package main

func main() {
	f := make(map[int][]string)
	f[1] = nil
	if _, ok := f[1]; ok {
		println("为 f 存入了一个值为nil的key")
	}
	delete(f, 1)

	if _, ok := f[1]; !ok {
		println("删除该key")
	}
}
