package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	a.SetString("3486784401", 10)
	b := new(big.Int)
	b.SetString("2621123515", 10)

	calculate(a, b)
}

func calculate(a, b *big.Int) {
	fmt.Printf("a = %s\nb = %s\n", a, b)

	mul := new(big.Int).Mul(a, b)
	div := new(big.Int).Div(new(big.Int).Set(a), b)
	sub := new(big.Int).Sub(a, b)
	add := new(big.Int).Add(a, b)

	fmt.Printf("a * b = %s\n", mul)
	fmt.Printf("a / b = %s\n", div)
	fmt.Printf("a - b = %s\n", sub)
	fmt.Printf("a + b = %s\n", add)
}
