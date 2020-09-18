package main

import (
	"fmt"
	"learnGo/makeStack/stack"
)

var fnc stack.Stack

func main()  {
	fnc.Push("anjay")
	fnc.Push(200)
	fnc.Push(20.20)
	fnc.Push([]string{"PHP", "GOLANG", "PYTHON", "C++"})

	for {
		item, err := fnc.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}