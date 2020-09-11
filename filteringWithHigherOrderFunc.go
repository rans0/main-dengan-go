package main
import "fmt"

func filter(number []int, check func(int) bool) []int {
	var even[] int
	for _, v := range number {
		if check(v) {
			even = append(even, v)
		}
	}
	return even
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func main() {
	arr := []int{1,2,3,4,5,6,7}
	fmt.Println(filter(arr, isEven))
}
