package logic

import (
	"geim/internal/logic/conf"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

func InitRPC(conf *conf.Rpc, object *Logic) {
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

type Reply struct {
	Id   string
	Room string
}

func (logic *Logic) CheckToken(token string, reply *Reply) error {
	//验证token
	hash, err := logic.redis.get(token)
	if err != nil {
		return err
	}
	reply.Id = hash["id"]
	reply.Room = hash["room"]
	log.Print("checktoken: id:", reply.Id, " room", reply.Room)
	return nil
}
