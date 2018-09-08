// Other than the Once and WaitGroup types, most are intended for use by low-level library routines.
// Higher-level synchronization is better done via channels and communication.
package main

import "sync"

func myfunc(wg *sync.WaitGroup, data int) {
	defer wg.Done()
	println(data)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
	}

	for i := 0; i < 100; i++ {
		go myfunc(&wg, i)
	}

	/* why dead lock, you can not pass the wg as value */
	wg.Wait()
}
