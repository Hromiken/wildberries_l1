package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	slice := create()
	target := rand.Intn(len(slice))
	sort.Ints(slice)
	start := time.Now()
	answer := binarySearch(slice, target)
	duration := time.Since(start)
	fmt.Printf("Target is number %d behind 0 and %d\n", target, len(slice)-1)
	fmt.Printf("Sucsess = %v\n", answer)
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

func binarySearch(slice []int, target int) bool {
	low, high := 0, len(slice)-1
	for low <= high {
		mid := (low + high) / 2
		if slice[mid] == target {
			return true
		}
		if slice[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
