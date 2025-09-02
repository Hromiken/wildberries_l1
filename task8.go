package main

import "fmt"

func main() {
	var number int64 = 5
	var position uint = 0
	bitValue := 0

	number = set(number, position, bitValue)
	fmt.Println(number)
}

func set(number int64, position uint, bitValue int) int64 {
	if bitValue == 1 {
		return number | (int64(1) << position)
	}
	return number &^ (int64(1) << position)
}
