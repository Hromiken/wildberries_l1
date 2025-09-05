package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "golang"
	num := 1234
	statment := false
	ch1 := make(chan int)
	ch2 := make(chan string)
	fmt.Println(variableType(str))
	fmt.Println(variableType(num))
	fmt.Println(variableType(statment))
	fmt.Println(variableType(ch1))
	fmt.Println(variableType(ch2))
}

func variableType(value interface{}) (answer string) {
	switch value.(type) {
	case int, string, bool:
		return fmt.Sprintf("type of variable = %T", value)
	}

	t := reflect.TypeOf(value)
	if t.Kind() == reflect.Chan {
		return fmt.Sprintf("type of variable = channel (%s)", t.String())
	}

	return fmt.Sprint("Unknown type of variable")
}
