package comet

import (
	"geim/internal/comet/conf"
	"net/rpc"
)

type Comet struct {
	conf        *conf.Comet
	Room_num    uint32
	channel_num uint32
	logicserver *rpc.Client
	rooms       map[string]*room
}

//初始化comet
func InitComet(conf *conf.Comet) (comet *Comet) {
	//初始化bucket
	comet = new(Comet)
	comet.conf = conf
	comet.rooms = make(map[string]*room, conf.Room_size)
	
	return

}
