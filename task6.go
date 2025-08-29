package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Выход по условию
	stopFlag := false
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if stopFlag {
				fmt.Println("Горутина завершает работу по условию")
				return
			}
			fmt.Println("Горутина (по условию) работает...")
			time.Sleep(500 * time.Millisecond)
		}
	}()
	time.Sleep(2 * time.Second)
	stopFlag = true

	//  Через канал уведомления
	done := make(chan os.Signal, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				fmt.Println("Горутина завершает работу по каналу")
				return
			default:
				fmt.Println("Горутина (канал) работает...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	close(done)

	// Через контекст
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина завершает работу по контексту")
				return
			default:
				fmt.Println("Горутина (контекст) работает...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)
	time.Sleep(2 * time.Second)
	cancel()

	// 4️⃣ runtime.Goexit()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Горутина вызываем Goexit")
		runtime.Goexit()
		fmt.Println("Это сообщение не будет напечатано")
	}()

	// Закрытие канала с данными
	job := make(chan int, 5)
	wg.Add(1)
	go func(ch chan int) {
		defer wg.Done()
		for v := range ch {
			fmt.Println("Горутина получила данные:", v)
		}
		fmt.Println("Горутина завершает работу по закрытию канала")
	}(job)

	for i := 0; i < 3; i++ {
		job <- i
	}
	close(job)

	// Паника / recover
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Горутина поймала панику и завершилась")
			}
		}()
		fmt.Println("Горутина вызывает панику")
		panic("аварийная ситуация")
	}()

	wg.Wait()
	fmt.Println("Все горутины завершены")
}
