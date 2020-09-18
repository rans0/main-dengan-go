package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

var nrchar, nrword, nrline int

func main() {
	nrchar, nrword, nrline = 0, 0, 0
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input, type S for stop program")

	for {
		input, err := inputReader.ReadString('\n')

		if err != nil {
			fmt.Printf("an error occured: %s\n", err)
			return
		}
		if input == "S\r\n" {
			fmt.Println("Here are the count")
			fmt.Printf("Number of character is %d\n", nrchar)
			fmt.Printf("Number of words is %d\n", nrword)
			fmt.Printf("Number of line is %d\n", nrline)
			os.Exit(0)
		}
		Counter(input)
	}
}

func Counter(input string) {
	nrchar += len(input) - 2
	nrword += strings.Count(input, " ") + 1
	nrline++
}
