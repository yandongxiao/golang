package ttt

import "testing"
import "fmt"

func TestMain(m *testing.M) {
    // 整个单元测试开始时执行
	fmt.Println("begin")
    // 整个单元测试结束时执行
	defer fmt.Println("end")
	m.Run()
}

func TestFoo(t *testing.T) {
    fmt.Println("testFoo stared")
    defer fmt.Println("testFoo ended")
    for i := 0; i<2; i++ {
        t.Run(fmt.Sprintf("%v", i), func(t *testing.T){
            fmt.Println("subtest stared")
            t.Cleanup(func() {
                fmt.Println("clean up")
            })
            fmt.Printf("subtest: %v\n", i)
        })
    }
}
