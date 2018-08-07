// Structs are composite types
// golang不存在继承, 数据和方法是完全解耦的关系
package main

import "fmt"

type Empty interface{}

type Engine struct {
	engine string
	size   int
}

type Car struct {
	Empty // 接口也可以匿名
	Engine
	name string
}

// 模拟继承
func main() {
	car := Car{
		Engine: Engine{
			engine: "1.2T",
			size:   100,
		},
		name: "丰田",
	}

	fmt.Println(car)
	fmt.Println(car.engine)
}
