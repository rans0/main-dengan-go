package main
import "fmt"

var fib [10]int64

func fibonacciWithArray() [10]int64 {

	fib[0] = 1
	fib[1] = 1

	for i:=2; i < 10; i++ {
		fib[i] = fib[i-1] + fib[i - 2]
	}
	return fib
}

func fibonacciWithSlice(term int) []int{

	arr := make([]int, term)
	arr[0], arr[1] = 1, 1

	for i:=2; i < term; i ++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr
}

func main() {

	arr := fibonacciWithArray()
	for i:=0; i < 10; i ++ {
		fmt.Printf( "The %d-th Fibonacci number using ARRAY is : %d\n", i, arr[i])
	}

	fmt.Println("##############################")

	slice := fibonacciWithSlice(10)
	for i, fib := range slice {
		fmt.Printf("The %d-th Fibonacci number using SLICE is : %d\n", i, fib)
	}
}