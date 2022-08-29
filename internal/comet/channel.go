package comet

import (
	"log"

	"github.com/gorilla/websocket"
)

type channel struct {
	conn *websocket.Conn
	id   string
	rm   *room
}

func (comet *Comet) newchannel(conn *websocket.Conn, userid, roomid string) {
	if conn == nil || userid == "" || roomid == "" {
		log.Printf("NewChannel error:userid:%v roomid:%v\n", userid, roomid)
		return
	}
	if _, ok := comet.rooms[roomid]; !ok {
		comet.newroom(roomid)
	}
	ch := channel{conn: conn, id: userid, rm: comet.rooms[roomid]}
	comet.rooms[roomid].channelmap[userid] = &ch
}
