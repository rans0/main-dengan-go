package main

import "fmt"

/*
Remark: When using type int the calculation is only correct up until 12!. This is, of course,
because an int can only contain integers that fit in 32 bits.
And, Go doesn’t usually warn against this overflow-error!
 */

func main() {
	for i := uint64(1); i <= 21; i++ {
		fmt.Println("Factorial of", i, " is", factorial(i))
	}
}

func factorial(num uint64) (fact uint64){

	if num <= 1 {
		return 1
	}

	fact = num * factorial(num-1)
	return
}