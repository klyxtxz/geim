package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func InitRPC(en *Engine) {
	rpc.Register(en)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1212")
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	http.Serve(l, nil)
}
