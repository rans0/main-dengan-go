package main

import "fmt"

var Days = map[int]string{
	1:"Sunday",
	2:"Monday",
	3:"Tuesday",
	4:"Wednesday",
	5:"Thursday",
	6:"Friday",
	7:"Saturday",
}

func findDays(n int) string {
	if value, key := Days[n]; key {
		return value
	} else {
		return "Wrong Key"
	}
}

func main()  {
	for i:=0; i <= 10; i++ {
		fmt.Printf("For key %d the value is %s\n", i, findDays(i))
	}
}
