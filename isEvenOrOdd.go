package main
import "fmt"

type flt func(int) bool

func isEven(num int) bool {
	if num % 2 == 0{
		return true
	}
	return false
}

func filter(sl[] int, f flt) (even, odd[] int) {
	for _, v := range sl{
		if f(v) {
			even = append(even, v)
		}else {
			odd = append(odd, v)
		}
	}
	return even, odd
}

func main() {
	var list[] int
	for i:=1; i <= 21; i ++ {
		list = append(list, i)
	}
	even, odd := filter(list, isEven)
	fmt.Println("List even :  ", even)
	fmt.Println("List odd : ", odd)
}
