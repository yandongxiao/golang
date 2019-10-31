// 继承是通过匿名内部类的方式实现的
// When an anonymous type is embedded in a struct,
// the visible methods of that type are embedded as well.
// embedded类型可以是struct，interface，alias type等
package main

import (
	"fmt"
)

type Engine interface {
	Start()
	Stop()
}

// 使用匿名内部类型来模拟面对对象编程中的继承.
// NOTE: 不但可以直接访问匿名类的field, 还可以直接访问匿名类的方法
// Golang规定：只要实现了接口定义的方法，即认为该类型实现了该interface。
type Car struct {
	wheelCount int
	Engine
}

type Mercedes struct {
	Car //anonymous field Car
}

// a behavior only available for the Mercedes
func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("Hi Angela!")
}

// define a behavior for Car
func (car Car) numberOfWheels() int {
	return car.wheelCount
}

// 即使Engine==nil, 因为*Car重载了它的方法，也OK。
// 即使*Car == nil, 因为Start和Stop方法没有对receiver进行解引用，也OK!
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

func ExampleMercedes() {
	m := Mercedes{Car{4, nil}} // Engine = nil
	fmt.Println(m.numberOfWheels())
	m.GoToWorkIn()
	m.sayHiToMerkel()

	// Output:
	// 4
	// Car is started
	// Car is stopped
	// Hi Angela!
}

func ExampleCar() {
	var car *Car
	car.GoToWorkIn()

	// Output:
	// Car is started
	// Car is stopped
}
