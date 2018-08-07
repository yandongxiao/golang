// magic.go
package main

import "fmt"

type Base struct{}

func (Base) Magic() { fmt.Print("base magic ") }

func (b Base) MoreMagic() {
	b.Magic() // must call Magic
	b.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() { fmt.Println("voodoo magic") }

func handle(b Base) {
	println("handle")
}

func main() {
	v := new(Voodoo)
	v.Magic()

	// Golang是静态的强类型的语言，且方法的本质是函数
	// NOTICE: Voodoo和Base是两个独立的类型，不能进行强制类型转换，更别说是隐士转换了）
	// cannot use *v (type Voodoo) as type Base in argument to handle)
	// handle(*v)

	// v.MoreMagic调用的本质是方法名寻找 + 传递对应field的过程
	// 所以，在MoreMagic内部调用的Magic方法时，MoreMagic函数拿到的形参是field，即Base类型
	// 所以，输出 base magic base magic
	v.MoreMagic()

}
