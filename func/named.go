package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println(MySqrt2(5))
}

// name the return variables - by default it will have 'zero-ed' values.
func MySqrt2(f float64) (ret float64, err error) {
	if f < 0 {
		//then you can use those variables in code
		ret = float64(math.NaN())
		err = errors.New("i won't be able to do a sqrt of negative number")
	} else {
		ret = math.Sqrt(f)
		//err is not assigned, so it gets default value nil
	}
	//automatically return the named return variables ret and err
	// 也可以使用return 1, nil. 强制输出其它结果
	return
}

func shadow() (val int) {
	// no new variables on left side of :=
	// 说明，returned named value 并没有处在一个更大的scope之中
	// val := 10

	if 1 == 1 { // correct
		// the Go 1 compilers disallow return statements without arguments
		// if any of the named return values is shadowed at the point of the return statement.
		// val is shadowed during return
		val := 10
		return val
	}

	if 1 == 1 { // compiler error
		val := 10
		return
	}

	return
}
