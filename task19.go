package main

import (
	"fmt"
)

func main() {
	example := "главрыба"
	answer := reverseString(example)
	fmt.Println(answer)
}

func reverseString(s string) string {

	r := []rune(s)

	for i, v := range s {
		r[i] = rune(v)
	}
	fmt.Println(string(r))
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	fmt.Println(string(r))
	return string(r)

}
