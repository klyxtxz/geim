package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type ClientInfo struct {
	Ch   string
	Room string
	Msg  []byte
}
type Recevier struct{}

func (r *Recevier) Handler(info *ClientInfo, reply *string) error {
	fmt.Print(*info)
	*reply = "ok"
	return nil
}

func main() {
	r := new(Recevier)
	rpc.Register(r)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1213")
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	http.Serve(l, nil)

}
