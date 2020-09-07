package main

import "fmt"

func ReturnMultiplevalue(num1, num2 int) (v1 int, v2 int, v3 int) {
	v1, v2, v3 = num1 + num2, num1 * num2, num1 - num2
	return
}

func main() {
	fmt.Println(ReturnMultiplevalue(20,10))
}