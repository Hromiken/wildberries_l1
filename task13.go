package main

func main() {
	num1 := 2
	num2 := 4

	changer(&num1, &num2)
}

func changer(num1, num2 *int) {
	*num1 = *num1 ^ *num2
	*num2 = *num1 ^ *num2
	*num1 = *num1 ^ *num2
}
