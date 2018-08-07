// time
// Note that the Go == operator compares not just the time instant but also the Location.
// 所以：time.Time = 数字 + 时区。
package main

import (
	"fmt"
	"time"
)

// NOTE: Location仅用于时区转化，而不对time内部的数据产生影响
func loadLocation() {
	println("loadLocation")
	// 等同于"UTC", 世界标准时间
	now := time.Now()
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

	fmt.Println(now) // Now的时区信息是Local
	fmt.Println(now.In(local1))
	fmt.Println(now.In(local2))
	fmt.Println(now.In(local3))
}

// Format
func format() {
	println("format")
	t := time.Now().UTC() // 此时返回了一个新的time.Time值，它是标准世界时间
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	// -0700 对应的输出是±hhmm，表示与世界标准时间的差值.
	// MST是比较有魔力的一个值，他会被转换为UTC（世界标准时间）
	fmt.Println(t.Format("Mon Jan 2 15:04:05 -0700 MST 2006")) // GMT和UTC是两种不同的时间记录方式，都是世界标准时间
	fmt.Println(t.Format("Mon Jan 2 15:04:05 GMT 2006"))       // GMT是一个普通的字符串

	now := time.Now() // 返回的是Local时区的时间
	// MST是比较有魔力的一个值，他会被转换为CST（北京时间）
	fmt.Println(now.Format("Mon Jan 2 15:04:05 -0700 MST 2006")) // GMT和UTC是两种不同的时间记录方式，都是世界标准时间
	fmt.Println(now.Format("Mon Jan 2 15:04:05 GMT 2006"))
}

func main() {
	loadLocation()
	format()
}
