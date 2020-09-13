package main

import "fmt"

type xStruct struct {
	flt float32
	int // anon field
	string // anon field
}

func main()  {
	ins1 := new(xStruct)
	(*ins1).flt = 1
	(*ins1).int = 2
	(*ins1).string = "insert using new"

	ins2 := xStruct{1, 2, "insert using literal expression"}

	fmt.Println(*ins1)
	fmt.Println(ins2)
	fmt.Println(ins1.flt, ins1.int, ins1.string)
}
