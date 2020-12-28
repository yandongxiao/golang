package main

import (
	"fmt"
)

type GPerson struct {
	Name string
}

// 与在外部声明单独声明Man, Woman相比，Ps将同一类型的数据聚合在了一起。
// 在最外层定义了匿名类. 外部访问方式：package.Ps.Man
// 与数组方式相比，它不需要下标，语义更清晰
// 与MAP方式相比，它不需要额外定义key的取值范围，更安全
var Ps struct {
	Man, Woman GPerson
}

func init() {
	Ps.Man.Name = "jack"
	Ps.Woman.Name = "lili"
}

func ExampleGourpedGlobals() {
	// 假设这是在另一个package中使用上面的全局变量
	fmt.Println(Ps.Man)
	fmt.Println(Ps.Woman)
	// Output:
	// {jack}
	// {lili}
}
