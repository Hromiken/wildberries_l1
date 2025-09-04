package main

import (
	"fmt"
	"sync"
)

func main() {
	workers := 3
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch1 := make(chan int)

	go func() {
		for _, v := range slice {
			ch1 <- v
		}
		close(ch1)
	}()

	ch2 := make(chan int)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(ch1, ch2, &wg, i)
	}
	go func() {
		wg.Wait()
		close(ch2)
	}()

	for v := range ch2 {
		fmt.Printf("Result Gouroutine = %d\n", v)
	}

}

func worker(ch1, ch2 chan int, wg *sync.WaitGroup, index int) {
	defer wg.Done()
	for v := range ch1 {
		ch2 <- v * 2
	}
}
