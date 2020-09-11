package main

import "fmt"

func enlarge(s []int, factor int) []int {

	newslice := make([]int, len(s)*factor)
	copy(newslice, s)
	return newslice
}

func enlargeAndFill(s []int, factor int) []int {
	large := len(s)*factor
	for i:=len(s); i < large; i++ {
		s = append(s, i)
	}
	return s
}



func main() {
	slice := []int{0,1,2}
	fmt.Println("Before enlarge is ", len(slice))
	fmt.Println(slice)
	slice = enlarge(slice, 5)
	fmt.Println("After enlarge is ", len(slice))
	fmt.Println(slice)
	slice2 := []int{0,1,2}
	slice2 = enlargeAndFill(slice2, 5)
	fmt.Println("After enlarge and fill is ", len(slice2))
	fmt.Println(slice2)
}
