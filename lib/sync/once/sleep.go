package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var once sync.Once

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Second * time.Duration(i))
			once.Do(func() {
				fmt.Println(i)
			})
		}(i)
	}

	select {}
}
