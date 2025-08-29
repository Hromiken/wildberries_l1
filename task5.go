package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	x := flag.Int("t", 3, "TimeLiving (seconds)")
	flag.Parse()
	var wg sync.WaitGroup

	job := make(chan int, 100)

	timer := time.NewTimer(time.Duration(*x) * time.Second)
	defer timer.Stop()

	workers := 2
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go timerWorkers(job, &wg, i)
	}

	data := 0
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-timer.C:
			fmt.Println("timer expired")
			close(job)
			wg.Wait()
			fmt.Println("Program is finished")
			return
		case <-ticker.C:
			job <- data
			data++
		}
	}
}

func timerWorkers(ch chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for v := range ch {
		log.Printf("Рабочий#%d получил:%d", i, v)
	}
	log.Printf("Рабочий #%d завершил работу", i)
}
