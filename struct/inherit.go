// inheritance_car.go
// When an anonymous type is embedded in a struct, the visible methods of that type are embedded as well.
// embedded类型可以是struct，interface，alias type
package main

import (
	"fmt"
)

type Engine interface {
	Start()
	Stop()
}

// 使用匿名内部类型来模拟面对对象编程中的继承. 不但可以直接访问匿名类的field, 还可以直接访问匿名类的方法
// Golang规定：只要实现了接口定义的方法，即认为该类型实现了该interface。没有implements或者extends等关键字
// 匿名内部类型 + 接口 即可实现多态！
// 内部类的本质是composite
type Car struct {
	wheelCount int
	Engine
}

// define a behavior for Car
func (car Car) numberOfWheels() int {
	return car.wheelCount
}

type Mercedes struct {
	Car //anonymous field Car
}

// a behavior only available for the Mercedes
func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("Hi Angela!")
}

func (car *Car) Start() {
	fmt.Println("Car is started")
}

func (car *Car) Stop() {
	fmt.Println("Car is stopped")
}

func (car *Car) GoToWorkIn() {
	// get in car
	car.Start()
	// drive to work
	car.Stop()
	// get out of car
}

func main() {
	m := Mercedes{Car{4, nil}}
	fmt.Println("A Mercedes has this many wheels: ", m.numberOfWheels())
	m.GoToWorkIn()
	m.sayHiToMerkel()
}
