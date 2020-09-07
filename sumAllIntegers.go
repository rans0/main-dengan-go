package main

import "fmt"

func main() {
	fmt.Println(sumInts(10,20,30,40))
	fmt.Println(sumInts(10,20,-30,5,2,5,1,18))
	fmt.Println(sumInts(-30,5,18))


}

func sumInts(num ...int) (sum int) {
	for len(num) == 0 {
		return
	}

	for _, v := range num {
		sum += v
	}

	return
}
