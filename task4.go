package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)


/*
Для завершения воркеров я использую context.WithCancel.
При получении сигнала Ctrl+C вызывается cancel(), и у контекста закрывается канал Done().
Все горутины-воркеры периодически проверяют этот канал и корректно выходят из цикла работы.
Такой способ удобен, потому что context — это стандартный механизм отмены в Go, он сразу распространяет сигнал на все горутины и позволяет аккуратно завершить программу.*/
func main() {
	workers := flag.Int("workers", 2, "Number of workers")
	flag.Parse()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-sig
		cancel()
	}()

	job := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go work(&wg, job, ctx, i)
	}

	data := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Получен сигнал, закрываем job...")
			close(job)
			wg.Wait()
			fmt.Println("Все воркеры завершены")
			return
		case job <- data:
			data++
			time.Sleep(1 * time.Second)
		}
	}

}

func work(wg *sync.WaitGroup, job chan int, ctx context.Context, i int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d: завершаюсь\n", i)
			return
		case data := <-job:
			fmt.Printf("worker %d получил задачу: %d\n", i, data)
		}
	}
}
