// time
// Note that the Go == operator compares not just the time instant but also the Location.
// 所以：time.Time = 数字 + 时区。
package main

import (
	"fmt"
	"testing"
	"time"
)

// NOTE: Location仅用于时区转化，而不对time内部的数据产生影响
func TestLoadLocation(t *testing.T) {
	// 已经做了时区转换
	// 2020-12-28 16:45:49.28318 +0800 CST m=+0.001168111
	now := time.Now()

	// 下面语句块的执行，没有修改 time.Now 的返回值
	// 2020-12-28 16:45:49.28318 +0800 CST m=+0.001168111
	local1, err1 := time.LoadLocation("") // 等同于传递"UTC"
	if err1 != nil {
		fmt.Println(err1)
	}

	// CST: 北京时间
	// 在国外获取中国北京时间，要用"PRC"
	local2, err2 := time.LoadLocation("Local")
	if err2 != nil {
		fmt.Println(err2)
	}

	// 美国洛杉矶时间
	local3, err3 := time.LoadLocation("America/Los_Angeles")
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Println(now) // Now 的时区信息是Local
	fmt.Println(now.In(local1))
	fmt.Println(now.In(local2))
	fmt.Println(now.In(local3))
}

// Format
func TestFormat(ts *testing.T) {
	t := time.Now().UTC() // 此时返回了一个新的time.Time值，它是标准世界时间
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	// -0700 对应的输出是±hhmm，表示与世界标准时间的差值.
	// MST是比较有魔力的一个值，他会被转换为UTC（世界标准时间）
	fmt.Println(t.Format("Mon Jan 2 15:04:05 -0700 MST 2006")) // 这个是UTC时间，GMT和UTC是两种不同的时间记录方式，都是世界标准时间
	fmt.Println(t.Format("Mon Jan 2 15:04:05 GMT 2006"))       // 这个是GMT时间，GMT是一个普通的字符串

	now := time.Now() // 返回的是Local时区的时间
	// MST是比较有魔力的一个值，他会被转换为CST（北京时间）
	fmt.Println(now.Format("Mon Jan 2 15:04:05 -0700 MST 2006")) // GMT和UTC是两种不同的时间记录方式，都是世界标准时间
	fmt.Println(now.Format("Mon Jan 2 15:04:05 GMT 2006"))       // 北京时间的GMT值，返回的是北京时间. 证明了GMT是一个普通字符串。
}
