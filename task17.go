package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	slice := create()
	sort.Ints(slice)

	target := rand.Intn(len(slice))

	start := time.Now()
	index := binarySearch(slice, target)
	duration := time.Since(start)

	fmt.Printf("Ищем число: %d\n", target)

	if index != -1 {
		fmt.Printf("Найден индекс: %d\n", index)
	} else {
		fmt.Println("Элемент не найден")
	}

	fmt.Printf("Time for binarySearch = %s\n", duration.String())
}

func create() []int {
	size := 50000
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(size)
	}
	return slice
}

func binarySearch(slice []int, target int) int {
	low, high := 0, len(slice)-1

	for low <= high {
		mid := (low + high) / 2

		if slice[mid] == target {
			return mid
		}

		if slice[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
