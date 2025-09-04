package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	fmt.Println(intersection(a, b))
}

func intersection(slice1, slice2 []int) (answer []int) {
	m := make(map[int]bool)
	for _, v := range slice1 {
		m[v] = true
	}

	seen := make(map[int]bool) // чтобы исключить дубликаты из slice2
	for _, v := range slice2 {
		if m[v] && !seen[v] {
			answer = append(answer, v)
			seen[v] = true
		}
	}
	return answer
}
