package main

import (
	"fmt"
	"strings"
)

func main() {
	test1 := checkUnique("abcd")
	test2 := checkUnique("abCdefAaf")
	test3 := checkUnique("aabcd")
	fmt.Println(test1, test2, test3)
}

func checkUnique(str string) bool {
	str = strings.ToLower(str)
	m := make(map[rune]int)
	for _, v := range str {
		m[v]++
	}
	for _, v := range m {
		if v > 1 {
			return false
		}
	}
	return true
}
