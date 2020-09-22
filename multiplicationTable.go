package main

import (
	"fmt"
	"sync"
	"time"
)

func printTable(n int, wg *sync.WaitGroup) {
	for i:=1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", i, n, i*n)
		time.Sleep(50 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for n:=2; n <= 12; n++ {
		wg.Add(1)
		go printTable(n, &wg)
	}
	wg.Wait()
}
