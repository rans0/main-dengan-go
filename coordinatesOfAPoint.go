package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64{
	return math.Sqrt(float64(p.x*p.x + p.y*p.y))
}

func (p *Point) Scale(flt float64) {
	p.x = p.x * flt
	p.y = p.y * flt
	return
}

func main() {
	n1 := Point{2, 5}
	fmt.Println(n1.Abs())
	n1.Scale(4)

	n2 := &Point{5, 2}
	fmt.Println(n2.Abs())
	n2.Scale(10)

}
