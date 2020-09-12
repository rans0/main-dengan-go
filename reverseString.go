package main

import "fmt"

func reverse(s string) string {
	arr := []byte(s)
	var rev [100]byte
	j := 0
	for i:=len(arr)-1; i >= 0; i-- {
		rev[j] = arr[i]
		j++
	}
	strev := string(rev[:len(arr)])
	return strev
}

func main() {
	arr := [...]string{"google", "checking", "login", "golang", "javascirpt"}
	for _, string := range arr {
		fmt.Println(reverse(string))
	}
}
