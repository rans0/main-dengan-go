package main

import "fmt"

func ReturnMultiplevalue(num1, num2 int) (v1, v2, v3, v4, v5 int) {
	v1, v2, v3, v4, v5 = num1 + num2, num1 * num2, num1 - num2, num1 / num2, num1 % num2
	return
}

func main() {

	sum, mult, min, div, mod := ReturnMultiplevalue(20,10)

	fmt.Println("sum : ", sum, "| mult : ", mult, "| min : ", min,
		"| div : ", div, "| mod : ", mod)
}