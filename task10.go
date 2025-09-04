package main

import (
	"fmt"
	"math"
)

func main() {
	slice := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(makeGroup(slice))
}

func makeGroup(slice []float64) map[int][]float64 {
	m := make(map[int][]float64)
	for _, v := range slice {
		group := int(math.Floor(v/10) * 10)
		m[group] = append(m[group], v)
	}
	return m
}
