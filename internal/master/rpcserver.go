package master

import (
	"geim/internal/master/conf"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

func InitRPC(conf *conf.Rpc, object *Master) {
	rpc.Register(object)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":"+strconv.Itoa(int(conf.Port)))
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	log.Print("RPC listening:", conf.Port)
	http.Serve(l, nil)
	defer l.Close()
}
