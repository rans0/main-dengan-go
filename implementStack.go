package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const LIMIT = 4

type Stack struct {
	ix int
	data [LIMIT]int
}

func (st *Stack) Push(n int) {
	if st.ix == LIMIT {
		return // stack full
	}
	st.data[st.ix] = n
	st.ix++ // increment
	return
}

func (st *Stack) Pop() int {
	if st.ix > 0 { // if stack not empty
		st.ix-- // decrease
		element := st.data[st.ix] // pop value
		st.data[st.ix] = 0 // after pop set to 0
		return element
	}
	return -1
}

func (st Stack) String() string {
	str := ""
	for ix:=0; ix < st.ix; ix++ {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(st.data[ix])  + "] "
	}
	return str
}

func main() {
	st := new(Stack)

	for ix:=0; ix < LIMIT; ix++ {
		n := rand.Intn(20)
		st.Push(n)
		fmt.Printf("%v\n", st)
	}

	for ix:=0; ix < LIMIT; ix++ {
		fmt.Printf("Popped : %d\n", st.Pop())
		fmt.Printf("%v\n", st)
	}

}
