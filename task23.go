package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Before:", slice, len(slice), cap(slice))

	slice = deleteElementByIndex(slice, 6)
	fmt.Println("After: ", slice, len(slice), cap(slice))
}

func deleteElementByIndex(slice []int, index int) []int {
	if len(slice) <= index || index < 0 {
		return slice
	}
	copy(slice[index:], slice[index+1:])
	slice[len(slice)-1] = 0
	return slice[:len(slice)-1]
}
