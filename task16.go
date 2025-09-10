package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := createSlice()
	start := time.Now()
	quickSort(s)
	duration := time.Since(start)
	fmt.Printf("Quick sort time = %s\n", duration)
}

func createSlice() []int {
	size := 50000
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(size)
	}
	return slice
}

func quickSort(slice []int) []int {
	// Базовый случай
	if len(slice) < 2 {
		return slice
	}
	// проверяем вдруг тут вообще одно действие
	if len(slice) == 2 {
		if slice[0] > slice[1] {
			slice[0], slice[1] = slice[1], slice[0]
			return slice
		}
	}

	left, right := 0, len(slice)-1 // индексы краев среза
	pivotIndex := len(slice) / 2   // центральный индекс
	pivot := slice[pivotIndex]     // элемент середины

	// Пока границы не сомкнутся
	for left <= right {
		// Пока элемент левой границы меньше элемента центра сжимаем слева
		for slice[left] < pivot {
			left++
		}
		// Пока элемент правой границы больше элемента сжимаем справа
		for slice[right] > pivot {
			right--
		}
		// Когда нашли пару «не на месте» (slice[left] => pivot, slice[right] <= pivot) — меняем их местами
		if left <= right {
			slice[left], slice[right] = slice[right], slice[left]
			left++
			right--
		}
	}

	//slice[:right+1] — сортируем левую часть (от начала до right включительно).
	if right > 0 {
		quickSort(slice[:right+1])
	}
	//slice[left+1:] — сортируем правую часть (от left+1 до конца включительно).
	if left < len(slice) {
		quickSort(slice[left:])
	}
	return slice
}
