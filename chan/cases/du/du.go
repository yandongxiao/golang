// package main does work like linux du command
// tow ways to exit:
//	1. user input content to standard input
//	2. all workers have finished theirs jobs
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// receive command from standard input to exited earlier
var end = make(chan struct{})

// limit the number of workers
// 为什么不使用 sync.WaitGroup 来限制worker的数量 ？
// 这是两种编程模式：
// sync.WaitGroup 将 worker 和协程进行强绑定，worker 只能在该协程内做任务。work 不断从 channel 中取任务，做任务。
// sema 模式下，同时只能有N个协程在做事情，具体是哪个协程，则不关心。
var sema = make(chan struct{}, 20)

func main() {
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = append(roots, ".")
	}

	go func() {
		_, _ = os.Stdin.Read(make([]byte, 1))
		close(end)
	}()

	// assign work to a single routine
	dataC := make(chan int64)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, dataC, &wg)
	}

	// 为什么要创建一个协程来做 wg.Wait 和 close(ch) ?
	// 主协程还有其它事情要做，比如监听end事件，监听worker的输出
	// 该协程的意义在于：所有 worker 都退出了，触发close(dataC)事件
	go func() {
		wg.Wait()
		close(dataC)
	}()

	// calculate
	var nfiles, nbytes int64
	tick := time.Tick(500 * time.Millisecond)
loop: // 在 for 和 select 同时使用的情况下，loop 是一种不错的选择
	for {
		select {
		case <-end: // user wants to exit earlier
			for range dataC {
				// 为什么要将dataC中的数据，读取完毕？
				// end 事件发生在前，worker会尽快推出，此处是，等待所有的worker都退出。
			}
			break loop
		case size, ok := <-dataC:
			if !ok { // all workers have finished theirs jobs
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			fmt.Printf("%d files  %v bytes\n", nfiles, nbytes)
		}
	}
	fmt.Printf("end: %d files  %v bytes\n", nfiles, nbytes)
}

func walkDir(root string, dataC chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case sema <- struct{}{}:
	case <-end:
		// 为什么需要监听 end ？
		return
	}

	defer func() { <-sema }()

	// 打开目录
	dir, err := os.Open(root)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	// 读取目录中所有的文件信息
	files, err := dir.Readdir(0)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		// NOTE: no return, because partial results may return
		// 虽然Readdir返回了错误，但是返回的部分，仍然可用
	}

	for _, file := range files {
		if file.IsDir() {
			wg.Add(1)
			dir := filepath.Join(root, file.Name())
			go walkDir(dir, dataC, wg)
		} else {
			dataC <- file.Size()
			time.Sleep(300 * time.Millisecond)
		}
	}
}
