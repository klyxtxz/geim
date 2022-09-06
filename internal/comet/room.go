package comet

import (
	"log"
)

type room struct {
	room_id      string
	channel_size uint32
	channelmap   map[string]*channel
}

func (comet *Comet) newroom(roomid string) {
	rm := new(room)
	rm.room_id = roomid
	rm.channel_size = 0
	rm.channelmap = make(map[string]*channel)
	comet.rooms[roomid] = rm
	
}

func (rm *room) push(msg, target string) {
	err := rm.channelmap[target].conn.WriteMessage(1, []byte(msg))
	if err != nil {
		log.Printf("push to %s error:%v\n", target, err)
		return
	}
	log.Printf("push to %s", target)
}
func (rm *room) pushall(msg, target string) {
	for chid, ch := range rm.channelmap {
		err := ch.conn.WriteMessage(1, []byte(msg))
		if err != nil {
			log.Printf("push to %s error:%v\n", chid, err)
			continue
		}
	}
	log.Printf("pushroom to %s", target)
}
