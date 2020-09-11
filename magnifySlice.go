package main

import "fmt"

func enlarge(s []int, factor int) []int{

	newslice := make([]int, len(s)*factor)
	copy(newslice, s)
	return newslice
}

func main() {
	slice := []int{0,1,2}
	fmt.Println("Before enlarge is ", len(slice))
	fmt.Println(slice)
	slice = enlarge(slice, 5)
	fmt.Println("After enlarge is ", len(slice))
	fmt.Println(slice)
}
