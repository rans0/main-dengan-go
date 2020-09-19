package main

import (
	"bufio"
	"fmt"
	"log"
	"io"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	title string
	price float64
	quantity int
}

func main() {
	bks := make([]Book, 1)
	file, err := os.Open("products.csv")
	if err != nil {
		log.Fatalf("Error %s opening file products.csv: ", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		// read one liner
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// remove \r and \n ( windows ) if linux only \n
		line = string(line[:len(line)-2])
		// remove separate delimeter
		sl := strings.Split(line, ";")
		book := new(Book)
		book.title = sl[0]

		book.price, err = strconv.ParseFloat(sl[1], 32)
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}

		book.quantity, err = strconv.Atoi(sl[2])
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}

		if bks[0].title == "" {
			bks[0] = *book
		}else {
			bks = append(bks, *book)
		}
	}
	fmt.Println("We have read the following books from the file: ")
	for _, bk := range bks {
		fmt.Println(bk)
	}
}
