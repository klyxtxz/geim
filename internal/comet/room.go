package comet

import "log"

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

func (rm *room) push(msg Msg) {
	rm.channelmap[msg.Target].conn.WriteMessage(1, msg.Content)
	log.Printf("push to %s", msg.Target)
}
func (rm *room) pushall(msg Msg) {
	for _, ch := range rm.channelmap {
		ch.conn.WriteMessage(1, msg.Content)
	}
	log.Printf("pushroom to %s", msg.Target)
}
