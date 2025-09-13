package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func main() {
	pos1 := NewPoint(1, 2)
	pos2 := NewPoint(4, 6)
	res := pos2.Distance(pos1)
	fmt.Printf("Расстояние между точками: %.2f\n", res)
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 {
	dx := other.x - p.x
	dy := other.y - p.y
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}
