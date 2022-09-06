package comet

import (
	"geim/internal/comet/conf"
	"log"
	"net/rpc"
)

type Comet struct {
	Addr         string
	conf         *conf.Comet
	Room_num     uint32
	channel_num  uint32
	logicserver  *rpc.Client
	masterserver *rpc.Client
	rooms        map[string]*room
}

//初始化comet
func InitComet(conf *conf.Comet) (comet *Comet) {
	//初始化bucket
	comet = new(Comet)
	comet.Addr = conf.Comet_id
	comet.conf = conf
	comet.rooms = make(map[string]*room, conf.Room_size)

	return

}

func (comet *Comet) Heartbeat() {
	type Comet struct {
		Addr  string
		Rooms []string
	}
	rms := []string{"aaa", "bbb", "ccc"}
	for rm, _ := range comet.rooms {
		rms = append(rms, rm)
	}
	arg := Comet{Addr: comet.Addr, Rooms: rms}
	rep := ""
	err := comet.masterserver.Call("Master.HeartBeat", &arg, &rep)
	if err != nil {
		log.Print("heartbeat err:", err)
		return
	}
	log.Print("heartbeat OK!")
}
