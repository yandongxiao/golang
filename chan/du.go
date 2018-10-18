package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var end = make(chan struct{})

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
		n.Wait()
		close(ch)
	}()

	// calculate
	tick := time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-end:
			for range ch {
			}
			break loop
		case size, ok := <-ch:
			if !ok {
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

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %v bytes\n", nfiles, nbytes)
}

var sema = make(chan struct{}, 20)

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
