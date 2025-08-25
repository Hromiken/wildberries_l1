package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	workers := flag.Int("w", 4, "Number of workers")
	flag.Parse()

	job := make(chan int, 100)

	var wg sync.WaitGroup
	for id := 1; id <= *workers; id++ {
		wg.Add(1)
		go work(&wg, job, id)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(job)
				return
			default:
				job <- rand.Intn(100)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	wg.Wait()
}
func work(wg *sync.WaitGroup, jobs <-chan int, id int) {
	defer wg.Done()
	for x := range jobs {
		fmt.Printf("Worker #%d send %d\n", id, x)
	}
}
