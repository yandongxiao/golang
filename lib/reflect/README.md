# reflect

## reflect 结构图

Method，SelectCase，SelectDir，SliceHeader，StringHeader，StructField，StructTag，
Type，Value等。关系如下：

![relation](./relation.png)

总结如下：

1. Type和Value是最重要的两个数据结构，尤其是Value
2. reflect.Value可以对内置类型，chan、map、struct等各种类型进行操作，如SetString, SetInt, Slice, MapKeys等。
3. reflect.Type的一大作用是给出了数据类型的模型，通过XxxOf()函数即可创建出该类型的一个新的实例。例如：```ArrayOf(count int, elem Type)```, ```FuncOf(in, out []Type, variadic bool) Type```, ```StructOf(fields []StructField) Type```等结构。

>StructOf用于动态构造struct类型对象。

## 常见的应用场景

应用场景：如果某个struct内部有interface{}类型的field？

解决办法：

1. 通过 type assertion 或者  type switch 来解决问题，这两种方式被称为 built-in reflection。
2. 需要通过reflect包来判断它有什么字段或者方法。

## Type Assertion

The convertibility for the first two are verified at compile time. The convertibility for the later two are verified at run time, by using a syntax called type assertion.

1. convert a non-interface value to an interface value, where the type of the non-interface value must implement the type of the interface value. 编译时判断conversion是否正确。
2. convert an interface value to an interface value, where the type of the source interface value must implement the type of the destination interface value. 编译时判断conversion是否正确。
3. convert an interface value to a non-interface value, where the type of the non-interface value must implement the type of the interface value.   // 编译通过的最基本要求，conversion不一定会成功，由运行时决定。In case of T is a non-interface type, if the dynamic type of i exists and is identical to T, then the assertion will succeed, otherwise, the assertion will fail.  When the assertion succeeds, the evaluation result of the assertion is a copy of the dynamic value of i. We can assertions of this case as value unboxing attempts.
4. convert an interface value to an interface value, where the type of the source interface value may or may not implement the type of the destination interface value. // 由运行时决定。

> 注意：In fact, for an interface value i with dynamic type as T, the method call i.m(...) is equivalent to the method call i.(T).m(...)

 所以，type assertion 在编译时通过，但是在运行时可能会出问题的。

在编译期对 i.(T) 的要求如下：

1. either a non-interface type which must implement the type of i
2. or an arbrtrary interface type.

在运行期：

1. 如果T是一个non-interface type，那么 i 的内部类型是可赋值给T的。
2. 如果T是一个interface type，那么要求 I 的内部类型，确实实现了T的接口规范。

Note, if a type assersion fails and the type assertion is used as a single-value expression (the second optional bool result is absent), a panic will occur.


## type-switch Control Flow Block

```go
switch aSimpleStatement; v := x.(type) {	// 括号内部是 type
	case TypeA:
		...
	case TypeB, TypeC:
		...
	case nil:   // 什么情况下会走到这个分支？
		...
	default:
		...
}
```

If the type denoted by a type name or type literal following a case keyword in a type-switch code block is not an interface type, then it must implement the interface type of the asserted value.

1. 对象x是一个interface类型。
2. 如果TypeA、TypeB不是一个interface类型，那么TypeA、TypeB一定要实现x规定的方法。
3. 如果TypeA、TypeB是一个interface类型，那么需要在运行时决定该case是否满足条件。

### case nil

```go
	var x interface{} = nil
	switch v := x.(type) {
	case nil:
		fmt.Printf("type of v is nil, %T\n", v)
	}
	
	output: type of v is nil, %v <nil>
```

### case default

```go
	a := 10
	var x interface{} = a
	switch v := x.(type) {
	default:
		fmt.Printf("type of x is %T\n", x)
		fmt.Printf("type of v is %T\n", v)
	}
	
	output: type of x is int
	output: type of v is int
```

1. 在上面的代码中，看不出v和x是相同的类型，都是interface{}类型。
2. %T 是可以获取到对象的本质类型的。

## package reflect

### 设置struct类型的某个字段。

```go
// CanSet reports whether the value of v can be changed.
// A Value can be changed only if it is addressable and was not
// obtained by the use of unexported struct fields.
// If CanSet returns false, calling Set or any type-specific
// setter (e.g., SetBool, SetInt) will panic.
func (v Value) CanSet() bool {
	return v.flag&(flagAddr|flagRO) == flagAddr
}
```

例子，定义 set 函数，对结构体的某个字段设置值。

```go
import (
	"fmt"
	"reflect"
)

func set(v interface{}, field string, newVal interface{}) {
	rv := reflect.ValueOf(v)
  if !rv.CanSet() {	// rv的底层类型是指针，rv本身不可以被Set，但是rv.Elem()可以
		rv = rv.Elem()
	}
  // 满足assignable关系
	rv.FieldByName(field).Set(reflect.ValueOf(newVal))	
}

type Person struct {
	Name string
	Age  int
}

func main() {
	var p Person
	set(&p, "Name", "jack")		// 必须传递P的指针
	set(&p, "Age", 10)
	fmt.Println(p)
}
```

1. 在调用set函数时，必须传递对象p的指针（否则编译出错），为什么？因为方法 `reflect.ValueOf` 的参数为 interface{}。
   1. 如果直接传递对象p，那么会造成对象p的复制。
   2. `Set`方法的修改，不会作用到对象p上
   3. Set 方法修改的对象，是一个你再也引用不到的对象。

### 调用对象的某个方法—参数以slice形式给出

```go
type S struct{}

func (s S) Add(a, b, c int) int {
	return a + b + c
}

// 调用对象object的方法
func call(object interface{}, methodName string, arguments ...interface{}) []reflect.Value {
	s := reflect.ValueOf(object)
  // NOTE: methodName的第一个字符必须是大写
	method := s.MethodByName(methodName)

	input := make([]reflect.Value, 0, len(arguments))
	for _, arg := range arguments {
		input = append(input, reflect.ValueOf(arg))
	}
	return method.Call(input)
}

func main() {
	s := S{}
	output := call(s, "Add", 1, 2, 3)
	for _, rt := range output {
		fmt.Println(rt.Interface())
	}
}
```

1. 通过反射调用方法，输入类型和输出类型都会被抽象为reflect.Value
2. 被调用的方法名称，一定要是exported method

### 调用对象的某个方法—参数以map的形式给出

```go
	s := reflect.ValueOf(S{})
	method := s.MethodByName("Add")

	x := 1
	y := 10
	// reflect.Type可以作为map的key
	m := make(map[reflect.Type]interface{})
	// y覆盖x
	m[reflect.TypeOf(x)] = x
	m[reflect.TypeOf(y)] = y

	input := make([]reflect.Value, 0)
	for i := 0; i < method.Type().NumIn(); i++ {
		input = append(input, reflect.ValueOf(m[method.Type().In(i)]))
	}

	out := method.Call(input)
	for i := 0; i < method.Type().NumOut(); i++ {
		fmt.Println(out[i].Interface())
	}
  
  // 20
```