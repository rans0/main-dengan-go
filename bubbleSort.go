package main

import "fmt"

func sortAsc(n []int) []int{
	for pass:=1; pass < len(n); pass++ {
		for i:=0; i < len(n) - pass; i++ {
			if n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i]
			}
		}
	}
	return n
}

func sortDesc(n []int) []int {
	for pass:=1; pass < len(n); pass++ {
		for i:=0; i < len(n) - pass; i++ {
			if n[i] < n[i+1] {
				n[i], n[i+1] = n[i+1], n[i]
			}
		}
	}
	return n
}

func main() {
	number := []int{4,5,6,1,2,3,2,1}
	fmt.Println("sorted ascending : ", sortAsc(number))
	fmt.Println("sorted descending : ", sortDesc(number))
}
