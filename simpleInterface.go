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

type NumberTwo struct {
	number int
}

func (n *NumberTwo) Set(param int) {
	n.number = param
}

func (n *NumberTwo) Get() int {
	return n.number
}

func f(call Call) int{
	switch call.(type) {
	case *Number :
		fmt.Println("Struct Number")
		call.Set(20)
		return call.Get()
	case *NumberTwo :
		fmt.Println("Struct Number Two")
		call.Set(10)
		return call.Get()
	default:
		fmt.Println("wrong")
	}
	return 0
}

func main() {
	var call1 Number
	fmt.Println(f(&call1))
	var call2 NumberTwo
	fmt.Println(f(&call2))
}