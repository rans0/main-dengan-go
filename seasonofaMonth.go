package main

import "fmt"

func Season(month int) string {

	switch month {
	case 1,2,12:
		return "Winter"
	case 3,4,5:
		return "Spring"
	case 6,7,8:Q
		return "Summer"
	case 9,10,11:
		return "Autumn"
	default:
		return "Unknown season"
	}

}

func main() {

	for i := 1; i <=13; i++ {
		month := i
		fmt.Println(Season(month))
	}

}
