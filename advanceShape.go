package main

import "fmt"


type Square struct {
	side float32
}

type Triangle struct {
	base, height float32
}

type AreaInterface interface {
	Area() float32
}

type PerimeterInterface interface {
	Perimeter() float32
}

func main() {
	var areaIntf AreaInterface
	var perimIntf PerimeterInterface

	sq := new(Square)
	sq.side = 5
	areaIntf = sq
	perimIntf = sq
	fmt.Printf("Area of Square is : %f\n", areaIntf.Area())
	fmt.Printf("Perimeter of Square is : %f\n", perimIntf.Perimeter())

	tr := new(Triangle)
	tr.height = 6
	tr.base = 5
	areaIntf = tr
	fmt.Printf("Area of Triangle is : %f\n", areaIntf.Area())
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (tr *Triangle) Area() float32 {
	return 0.5 * tr.base * tr.height
}

func (sq *Square) Perimeter() float32 {
	return 4 * sq.side
}


