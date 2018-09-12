package main

import (
	"flag"
	"fmt"
)

type Person struct {
	Name string
}

// 在最外层定义了匿名类. 访问方式：xxx.Ps.Man
// 与数组的方式相比，它不需要下标，语义更清晰
// 与在外部声明单独声明Man, Woman相比，Ps将同一类型的数据聚合在了一起
var Ps struct {
	Man, Woman Person
}

func init() {
	Ps.Man.Name = "jack"
	Ps.Woman.Name = "lili"
}

// 假设这是在另一个package中使用上面的全局变量
func main() {
	flag.Parse()

	// 如果是
	fmt.Println(Ps.Man)
	fmt.Println(Ps.Woman)
}
