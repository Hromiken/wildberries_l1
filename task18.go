package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var num int64
	var wg sync.WaitGroup

	numWorkers := 3

	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&num, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Результат:", num)
}
