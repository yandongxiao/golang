// NOTE: Except for special, low-level applications, synchronization is better done with
// channels or the facilities of the sync package.
// Share memory by communicating; don't communicate by sharing memory.
package main

import "fmt"

import "sync/atomic"

func main() {

	// SwapT: 等价于：old = *addr; *addr = new; return old
	v1 := int32(1)
	fmt.Println(atomic.SwapInt32(&v1, 2), v1)

	// CompareAndSwapT: 等价于:
	// if *addr == old {
	//		*addr = new
	//		return true
	// }
	// return false
	fmt.Println(atomic.CompareAndSwapInt32(&v1, 2, 4), v1)

	// AddT：等价于：*addr += delta; return *addr
	fmt.Println(atomic.AddInt32(&v1, 4), v1)

	// 相当于对值的读和写：LoadT and StoreT: "return *addr" and "*addr = val"
	atomic.StoreInt32(&v1, 16)
	fmt.Println(atomic.LoadInt32(&v1))

	// NOTE: 返回值与v1并非是同一个变量
	fmt.Println(atomic.AddInt32(&v1, 16))
}
