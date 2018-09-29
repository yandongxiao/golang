package main

import (
	"log"
	"os"
	"sync"
)

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	// sizes是否结束依赖wg.Wait
	// wg.Wait如果放在for range的前面，由于sizes是同步chan，导致死锁
	// wg.Wait如果放在for range的后面，则for循环不会退出
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
