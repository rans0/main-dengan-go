package main

import (
	"fmt"
	"sync"
)

func deposit(balance *int, amount int, myMutex *sync.Mutex, myWaitGroup *sync.WaitGroup) {
	myMutex.Lock()
	*balance += amount
	myMutex.Unlock()
	myWaitGroup.Done()
}

func withdraw(balance *int, amount int, myMutex *sync.Mutex, myWaitGroup *sync.WaitGroup) {
	myMutex.Lock()
	*balance -= amount
	myMutex.Unlock()
	myWaitGroup.Done()
}

func main() {
	balance := 100
	var myMutex sync.Mutex
	var myWaitGroup sync.WaitGroup

	myWaitGroup.Add(2)
	go deposit(&balance, 15, &myMutex, &myWaitGroup)
	withdraw(&balance, 70, &myMutex, &myWaitGroup)

	myWaitGroup.Wait()
	fmt.Println(balance)
}
