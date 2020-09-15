package main

import "fmt"

type Call interface {
	Set(int)
	Get()int
}

type Number struct {
	number int
}

func (n *Number) Set(param int) {
	n.number = param
}

func (n *Number) Get() int {
	return n.number
}

func f(call Call) int{
	call.Set(10)
	return call.Get()
}

func main() {
	var call Number
	fmt.Println("Before set : ", call.number)
	fmt.Println("After set : ", f(&call))
}