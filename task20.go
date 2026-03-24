package main

import (
	"fmt"
)

func main() {
	str := "snow dog sun"
	fmt.Println(reverseWords(str))
}

func reverseWords(s string) string {
	r := []rune(s)

	reverse(r, 0, len(r)-1)

	start := 0
	for i := 0; i <= len(r); i++ {
		if i == len(r) || r[i] == ' ' {
			reverse(r, start, i-1)
			start = i + 1
		}
	}

	return string(r)
}

func reverse(r []rune, i, j int) {
	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
}
