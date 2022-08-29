package comet

import (
	"errors"
	"geim/internal/comet/conf"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

type Msg struct {
	CallerId  string
	Operation int
	Target    string
	Content   []byte
}

func InitRPCServer(conf *conf.Rpc, object *Comet) {
	rpc.Register(object)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":"+strconv.Itoa(int(conf.Port)))
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	log.Print("RPC listening:", conf.Port)
	http.Serve(l, nil)
}

func (comet *Comet) PushId(msg Msg, reply *string) error {
	for _, room := range comet.rooms {

		if _, ok := room.channelmap[msg.Target]; ok {
			room.push(msg)
			rep := "success"
			reply = &rep
			return nil
		}
	}

	return errors.New("404 Not found")
}

func (comet *Comet) PushRoom(msg Msg, reply *string) error {
	comet.rooms[msg.Target].pushall(msg)
	rep := "success"
	reply = &rep
	return nil

}

func (comet *Comet) PushAll(msg Msg, reply *string) error {
	for _, rm := range comet.rooms {
		rm.pushall(msg)
	}
	rep := "success"
	reply = &rep
	return nil
}
