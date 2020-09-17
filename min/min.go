package main

import (
	"fmt"
	"learnGo/min/sort"
)

func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := sort.IntArray(data)
	m := sort.Min(a)
	fmt.Printf("The minimal value is %v\n", m)
}

func strs() {
	data := []string{"ddd", "eee", "bbb", "ccc", "aaa"}
	a := sort.StringArray(data)
	m := sort.Min(a)
	fmt.Printf("The minimal value is %v\n", m)
}

func main() {
	ints()
	strs()
}
