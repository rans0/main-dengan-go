package main

import "fmt"

func fibonacci(n int) chan int {
	myChannel := make(chan int)
	go func() {
		k:=0
		for i, j := 0, 1; k < n; k++ {
			myChannel <- i
			i, j = i+j, i
		}
		close(myChannel)
	}()
	return myChannel
}

func main() {
	for i:= range fibonacci(5) {
		fmt.Println(i)
	}
}