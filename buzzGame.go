package main

import (
	"fmt"
	"time"
	"math/rand"
)

/*
	if player buzz 1 buzz first
	output will :
	Player 1 buzzed
	Player 2 buzzed

	if player buzz 2 buzz first
	output will :
	Player 2 buzzed
	Player 1 buzzed
 */
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		channel1 <- "Player 1 buzzed"
	}()

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		channel2 <- "Player 2 buzzed"
	}()

	for i:=0; i < 2; i++ {
		select {
		case player1 := <-channel1:
			fmt.Println(player1)
		case player2 := <-channel2:
			fmt.Println(player2)
		}
	}

}
