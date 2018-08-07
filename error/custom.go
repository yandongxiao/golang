// 1. A package can also define its own specific Error with additional methods, like net.Error.
// 2. naming convention: Error types end in “Error” and error variables are called (or start with) “err” or “Err”.
// 3. In case different possible error-conditions can occur, it may be useful to test with a type assertion
//      or type switch for the exact error, and possibly try a remedy or a recovery of the error-situation
//		如果方法返回不同类型的err，还要让调用者判断，无疑是很重的负担
package main

import (
	"fmt"
	"strconv"
)

type UError uintptr

func (err UError) Error() string {
	return strconv.Itoa(int(err))
}

func main() {
	// interface的比较规则如下：
	//	1. 首先比较的两个对象必须是同一种类型，即实现了相同的接口
	//	2. 如果底层类型不相同，则返回false
	//	3. 如果底层类型相同，但是类型不支持比较，如map, slice, 则报错
	//	4. 如果支持比较，则按照底层类型的比较规则进行比较.
	err := UError(3)
	if err == 3 {
		fmt.Println("you can compare with interface")
	}
}
