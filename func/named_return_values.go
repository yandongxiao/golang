package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println(MySqrt2(5))
}

// name the return variables - by default it will have 'zero-ed' values i.e. numbers are 0, string is empty, etc.
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
