// 假如我们在写一个golang库-A. 如果该库在调用其它库-B函数时，
// 发生panic(这个是不会在文档中说明的)，该如何自处?
// 1. 调用recover函数，恢复一切错误. 将panic转换为error;
// 2. 但是, 你不应该在库中的每个export函数都设置defer + recover.
// golang的标准库函数也没有这么做. 所以，方法1是不现实的.
// 3. 更加可取的方法，recover的错误一定是库A的panic。如果是库B的panic，
// 即使在库A的defer语句中捕捉到了，也继续向上抛
// 以下是实现思路如下（将它考虑成一个package）

package gerror

import (
	"fmt"
	"os"
)

// 库A的所有panic的值必须是Error类型
type Error string

func (e Error) Error() string {
	return string(e)
}

// NOTICE: package内部可以通过调用该函数进行panic. 类似Java的异常
func newError(err string) {
	panic(Error(err))
}

// repanic会将两个panic的信息都打印出来的
// panic: open /tmp/xxyy: no such file or directory [recovered]
// panic: interface conversion: interface {} is *os.PathError,
// not main.Error]
func Compile() (num int, err error) {
	// export函数使用defer+recover对所有panic进行捕获
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			num = 0         // NOTE: defer内部可以修改返回值
			err = e.(Error) // 过滤库A的panic. 其它库的panic会因为type assertion, repanic
		}
	}()

	if _, err = os.Open("/tmp/xxyy"); err != nil {
		panic(err)
	}
	return 0, nil
}

func ExampleRepanic() {
	defer func() {
		fmt.Println(recover())
	}()
	Compile()
	// Output:
	// open /tmp/xxyy: no such file or directory
	// interface conversion: interface {} is *os.PathError, not main.Error
}
