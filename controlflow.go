package main

import (
	"fmt"
	"time"
)

func main() {

	// if else
	fmt.Println("======IF ELSE======")

	input := 8

	if input < 0 {
		fmt.Println("Input tidak boleh negatif")
	} else if input%2 == 0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}

	//switch
	fmt.Println("======SWITCH======")

	hariLahir := time.Friday

	switch hariLahir {
	case time.Monday:
		fmt.Println("Hari ini senin")
	case time.Tuesday:
		fmt.Println("Hari selasa")
	default:
		fmt.Println("input tidak valid")
	}

	//for
	fmt.Println("======FOR======")

	for index := 4; index >= 2; index-- {
		fmt.Println(index)
	}

	//for range
	fmt.Println("======FOR RANGE======")

	numbers := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	for i, v := range numbers {
		fmt.Println("Array Element", i, "is", v)
	}

	//break
	fmt.Println("======BREAK======")

	for index := 0; index < 5; index++ {
		if index == 3 {
			break
		}
		fmt.Println(index)
	}

	//continue
	fmt.Println("======CONTINUE======")

	for index := 0; index < 5; index++ {
		if index == 3 {
			continue
		}
		fmt.Println(index)
	}

}
