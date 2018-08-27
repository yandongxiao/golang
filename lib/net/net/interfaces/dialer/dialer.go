package main

import (
	"fmt"
	"net"
	"time"

	"github.com/felixge/tcpkeepalive"
)

func main() {
	dialer := new(net.Dialer)

	// 超时设置
	// Timeout is the maximum amount of time a dial will wait for a connect to complete.
	// If Deadline is also set, it may fail earlier.
	// The default is no timeout.
	// NOTE:
	// With or without a timeout, the operating system **may** impose its own earlier timeout.
	// For instance, TCP timeouts are often around 3 minutes.
	dialer.Timeout = time.Second
	dialer.Deadline = time.Now().Add(time.Minute) // 属性与Timeout类似

	// 设置本地地址
	dialer.LocalAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9876")

	// 设置TCP的KeepAlive属性，参见: https://www.cnblogs.com/havenshen/p/3850167.html
	// NOTICE: 注意TCP中KeepAlive与http中keep-alive之间的区别
	// 在mac下，查看KeepAlive相关的三个属性：sysctl net.inet.tcp | grep keep
	//
	// SetKeepAlivePeriod However, this method currently behaves different for different operating systems.
	// On OSX, it will modify the idle time before probes are being sent.
	// On Linux however, it will modify both the idle time, as well as the interval that probes are sent at.
	// So calling SetKeepAlivePeriod with an argument of 30 seconds will cause a total timeout of 10 minutes
	// and 30 seconds for OSX (30 + 8 * 75), but 4 minutes and 30 seconds on Linux (30 + 8 * 30).
	// net.inet.tcp.keepcnt: 8 没有被修改的接口
	//
	// If zero, keep-alives are not enabled. 即，不管连接处于断开还是连接状态，只管进行IO操作
	dialer.KeepAlive = 0

	conn, _ := dialer.Dial("tcp", "127.0.0.1:8888")
	fmt.Printf("client: local addr=%#v, remote addr=%#v, %T", conn.LocalAddr().String(), conn.RemoteAddr().String(), conn)
	// tcpConn := conn.(*net.TCPConn)
	// tcpkeepalive 是一个第三方工具包，它可以被用来分别设置KeepAlive的三个属性
	tcpConn, _ := tcpkeepalive.EnableKeepAlive(conn)
	tcpConn.SetKeepAlive(true)
	tcpConn.SetKeepAliveIdle(time.Second)
	tcpConn.SetKeepAliveCount(1)
	tcpConn.SetKeepAliveInterval(time.Second)

	// KeepAlive模拟
	// 首先因为这是TCP协议的属性，在应用层是无法模拟KeepAlive的行为的
	// 其次，通过杀死Server的方式，无论KeepAlive是否设置，Read操作都能立刻感受到服务端被关闭。
	//       不是说进程被杀死了，进程的所有连接就不再通信了，它们应该还是走TCP的断开协议的
	// 最后, 最好是通过网络来模拟KeepAlive的作用
	time.Sleep(time.Second * 5)
	data := make([]byte, 1024)
	fmt.Println(conn.Read(data))
}
