package main

import (
	"fmt"
	"math"
)

func main() {
	slice := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10.0

	grouped := makeGroup(slice, step)
	fmt.Println(grouped)
}

func makeGroup(slice []float64, step float64) map[int][]float64 {
	m := make(map[int][]float64)
	for _, v := range slice {
		group := int(math.Floor(v/step) * step)
		if v < 0 && float64(group) != v {
			group += int(step)
		}
		m[group] = append(m[group], v)
	}
	return m
}
