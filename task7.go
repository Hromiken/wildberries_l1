package main

import (
	"log"
	"sync"
)

/*Реализовать конкуретную запись в map'у*/
func main() {
	m := make(map[int]int)
	var m2 sync.Map

	ch := make(chan int, 100)
	ch2 := make(chan int, 100)

	var mu sync.Mutex
	var wg, wg2 sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go mapWriter(m, &wg, ch, &mu)
	}
	for i := 0; i < 2; i++ {
		wg2.Add(1)
		go mapWriter2(&m2, &wg2, ch2)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	go func() {
		wg2.Wait()
		close(ch2)
	}()

	for v := range ch {
		log.Println(v)
	}
	for value := range ch2 {
		log.Println(value)
	}
}

func mapWriter(m map[int]int, wg *sync.WaitGroup, ch chan int, mu *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		mu.Lock()
		m[i] = i * i
		v := m[i]
		mu.Unlock()
		ch <- v
	}
}

func mapWriter2(m *sync.Map, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for i := 200; i < 300; i++ {
		m.Store(i, i*i)
		v, _ := m.Load(i)
		ch <- v.(int)
	}
}
