package main

import "fmt"

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(sub(slice))
}

func sub(slice []string) (answer []string) {
	m := make(map[string]struct{})

	for _, v := range slice {
		m[v] = struct{}{}
	}

	for k := range m {
		answer = append(answer, k)
	}

	return answer
}
