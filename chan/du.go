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

// recevie command from standard input to exited earlier
var end = make(chan struct{})

// limit the number of workers
var sema = make(chan struct{}, 20)

func main() {
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = append(roots, ".")
	}

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(end)
	}()

	// assign work to a single routine
	ch := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, ch, &n)
	}
	go func() {
		n.Wait() // NOTE: means all workers have exited
		close(ch)
	}()

	// calculate
	tick := time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-end: // user wants to exit earlier
			for range ch { // make sure all workers have exited
			}
			break loop
		case size, ok := <-ch:
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

func walkDir(root string, ch chan<- int64, n *sync.WaitGroup) {
	defer n.Done()

	select {
	case sema <- struct{}{}:
	case <-end:
		return
	}

	defer func() { <-sema }()

	dir, err := os.Open(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	files, err := dir.Readdir(0)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		// NOTE: no return, because partial results may return
	}

	for _, file := range files {
		if file.IsDir() {
			n.Add(1)
			dir := filepath.Join(root, file.Name())
			go walkDir(dir, ch, n)
		} else {
			ch <- file.Size()
			time.Sleep(300 * time.Millisecond)
		}
	}
}
