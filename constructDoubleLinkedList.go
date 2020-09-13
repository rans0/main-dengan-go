package main

import (
	"fmt"
	"container/list"
)

func insertListElement(n int)(*list.List) {
	lst := list.New()
	for i:=1; i<=n; i++ {
		lst.PushBack(i)
	}
	return lst
}

func main() {
	n := 5
	myList := insertListElement(n)
	for e:=myList.Front(); e!=nil; e = e.Next(){
		fmt.Println(e.Value)
	}
}