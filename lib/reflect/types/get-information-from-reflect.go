package main

import (
	"fmt"
	"reflect"
)

func inspect(d interface{}) {
	fmt.Println("--------------------------------")

	var (
		dt reflect.Type
		dv reflect.Value
	)

	dv = reflect.ValueOf(d)
	fmt.Printf("Value: %v\n", dv.Interface())

	dt = reflect.TypeOf(d)
	fmt.Printf("Type: %s\n", dt)
	fmt.Printf("Name: %s\n", dt.Name())
	fmt.Printf("Kind: %s\n", dt.Kind())
	fmt.Println()

	if dt.Kind() == reflect.Struct {
		for i := 0; i < dv.NumField(); i++ {
			fmt.Printf("Field %d\n", i)
			ft := dt.Field(i)
			fv := dv.Field(i)

			if !ft.Anonymous {
				fmt.Printf("Type: %s\n", ft.Type.Name())
				fmt.Printf("Kind: %s\n", ft.Type.Kind())
				fmt.Printf("Name: %s\n", ft.Name)
			}

			fmt.Printf("Value: %v\n\n", fv.Interface())
		}

		for i := 0; i < dv.NumMethod(); i++ {
			fmt.Printf("Method %d\n", i)
			mt := dt.Method(i)
			fmt.Printf("Type: %s\n", mt.Type.Name())
			fmt.Printf("Kind: %s\n", mt.Type.Kind())
			fmt.Printf("Name: %s\n", mt.Name)
			fmt.Println()
		}
	} else if dt.Kind() == reflect.Func {
		for i := 0; i < dt.NumIn(); i++ {
			pt := dt.In(i)
			fmt.Printf("In %d\n", i)
			fmt.Printf("Type: %s\n", pt.Name())
			fmt.Printf("Kind: %s\n", pt.Kind())
			fmt.Println()
		}

		for i := 0; i < dt.NumOut(); i++ {
			pt := dt.Out(i)
			fmt.Printf("Out %d\n", i)
			fmt.Printf("Type: %s\n", pt.Name())
			fmt.Printf("Kind: %s\n", pt.Kind())
			fmt.Println()
		}
	}

	fmt.Println("--------------------------------")
}

/**
 * Test types.
 */
type ErdosNumber uint

type Person struct {
	Name string
	Age  int
	E    ErdosNumber
}

/**
 * Test functions.
 */
func (p Person) SayHello() string {
	return fmt.Sprintf("Hello, %s!", p.Name)
}

func CheckErdos(e ErdosNumber) (string, error) {
	value := uint(e)
	if value == 0 {
		return "", fmt.Errorf("Invalid erdos number")
	}

	return fmt.Sprintf("Your erdos number is %d.", value), nil
}

func Add(a, b int) int {
	return a + b
}

/**
 * Application entry point.
 */
func main() {
	// inspect a struct
	alex := Person{Name: "alex", Age: 22, E: ErdosNumber(2)}
	inspect(alex)

	// create a struct
	dv := reflect.New(reflect.TypeOf(Person{}))
	dve := dv.Elem()
	dve.FieldByName("Name").SetString("alex")
	dve.FieldByName("Age").SetInt(22)

	ev := reflect.New(dve.FieldByName("E").Type())
	eve := ev.Elem()
	eve.SetUint(2)
	dve.FieldByName("E").Set(eve)
	inspect(dve.Interface())

	// inspect functions
	inspect(CheckErdos)
	inspect(Add)

	// call a function
	av := reflect.New(reflect.TypeOf(int(1)))
	ave := av.Elem()
	ave.SetInt(2)

	bv := reflect.New(reflect.TypeOf(int(2)))
	bve := bv.Elem()
	bve.SetInt(3)

	addv := reflect.ValueOf(Add)
	res := addv.Call([]reflect.Value{ave, bve})
	fmt.Printf("Add(2, 3) => ")
	fmt.Println(res[0].Interface())
}
