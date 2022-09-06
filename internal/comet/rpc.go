package comet

import (
	"geim/internal/comet/conf"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

func InitRPCServer(conf *conf.Rpc, object *Comet) {
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

type Pushid struct {
	Target  string
	Room    string
	Content string
}

func (comet *Comet) PushId(msg Pushid, reply *string) error {
	comet.rooms[msg.Room].push(msg.Content, msg.Target)

	return nil
}

type Pushroom struct {
	Target  string
	Content string
}

func (comet *Comet) PushRoom(msg Pushroom, reply *string) error {
	comet.rooms[msg.Target].pushall(msg.Content, msg.Target)
	*reply = "success"

	return nil

}

func (comet *Comet) PushAll(msg string, reply *string) error {
	for rmid, rm := range comet.rooms {
		rm.pushall(msg, rmid)
	}
	*reply = "success"

	return nil
}
