package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func Callcomet(msg []byte) {

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1212")
	if err != nil {
		log.Fatal("dialing: ", err)
	}
	var reply string
	err = client.Call("Engine.BroadCast", string(msg), &reply)
	if err != nil {
		log.Fatal("rpc error: ", err)
	}
	fmt.Printf("reply: %v", reply)

}
