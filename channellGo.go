package main

import "fmt"

func sendValues(myIntChannel chan int) {
	for i:=0; i < 5; i++ {
		myIntChannel <- i // sending value
	}
	close(myIntChannel)
}

func sendValues2(myIntChannel2 chan int) {
	for i:=0; i < 5; i++ {
		myIntChannel2 <- i // sending value
	}
}

func main() {
	myIntChannel := make(chan int)
	go sendValues(myIntChannel)
	for value := range myIntChannel {
		fmt.Println(value)
	}

	myIntChannel2 := make(chan int)
	defer close(myIntChannel2)
	go sendValues2(myIntChannel2)
	for i:=0; i < 5; i++ {
		fmt.Println(<-myIntChannel2)
	}

}
