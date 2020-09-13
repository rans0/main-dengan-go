package main

import "fmt"

type Rectangle struct {
	length, width int
}

func (a *Rectangle) Area() int {
	return a.length * a.width
}

func (p *Rectangle) Peremiter() int {
	return 2 * (p.length + p.width)
}

func main() {
	n1 := Rectangle{10, 2}
	fmt.Println(n1.Area())
	fmt.Println(n1.width, n1.length)

	n2 := Rectangle{7, 4}
	fmt.Println(n2.Peremiter())
	fmt.Println(n2)
}
