package main

import (
	"fmt"
	"sync"
)

func main() {
	slice := []int{2, 4, 6, 8, 10}

	jobs := make(chan int, len(slice))
	results := make(chan int, len(slice))

	workers := 3
	var wg sync.WaitGroup

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go square(i, jobs, results, &wg)
	}

	for _, v := range slice {
		jobs <- v
	}
	close(jobs)

	wg.Wait()
	close(results)

	for res := range results {
		fmt.Println(res)
	}
}

func square(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs {
		value := n * n
		results <- value
		fmt.Printf("Value %d has been squared to %d\n", n, value)
	}
}
