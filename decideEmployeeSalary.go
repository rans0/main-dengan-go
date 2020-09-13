package main

import "fmt"

type Employee struct {
	salary float32
}

func (this *Employee) giveRaise(pct float32) {
	this.salary += this.salary * pct
	return
}

func main() {
	emp := Employee{1000}
	emp.giveRaise(0.5)
	fmt.Println(emp.salary)
}
