package main

import (
	"fmt"
	"log"
	"net/rpc"
	"learnGo/connRPCinGO/rpcobjects"
)

const serverAddress = "localhost"

func main() {
	client, err := rpc.Dial("tcp", serverAddress + ":3001")
	if err != nil {
		log.Fatal("Error dialing: ", err)
	}
	// Synchronous call
	args := &rpcobjects.Args{3, 6}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
}
