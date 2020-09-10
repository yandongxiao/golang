package main

func main() {
	var x *[9]int
	x = &[...]int{8: 0}
	y := x[1:3]
	_ = y[5:7] // line 7
	_ = y[5:]  // line 8, 为什么这一行会崩溃？y[5:] == y[5:len(y):cap(y)] 违反了0<=low<=high<=cap的原则
	x = nil
	for _, _ = range x {
	} // line 10
	_ = len(y) / len(x) // line 11
}
