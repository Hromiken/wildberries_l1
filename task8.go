package main

import "fmt"

func main() {
	var number int64 = 5
	var position uint = 1
	bitValue := 0

	result, err := set(number, position, bitValue)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println(result)
}

func set(number int64, position uint, bitValue int) (int64, error) {
	if bitValue != 0 && bitValue != 1 {
		return 0, fmt.Errorf("бит может быть только 0 или 1")
	}

	if position == 0 || position > 64 {
		return 0, fmt.Errorf("позиция должна быть от 1 до 64")
	}

	pos := position - 1

	if bitValue == 1 {
		return number | (int64(1) << pos), nil
	}

	return number &^ (int64(1) << pos), nil
}
