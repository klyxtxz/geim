package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/websocket"
)

type Channel struct {
	conn *websocket.Conn
	id   string
	rm   *Room
}

//新建一个频道
func NewChannel(conn *websocket.Conn, id string) *Channel {
	return &Channel{conn: conn, id: id}
}

//频道开始接收消息
func (chn *Channel) Serve() {
	//退出时删除对象引用，将指针设为空，等待GC回收。
	defer chn.rm.DelCh(chn.id)
	for {

		msgType, msg, err := chn.conn.ReadMessage()
		if err != nil {
			if strings.Contains(err.Error(), "close") {
				log.Printf("连接关闭,ID:%v", chn.id)
				return
			}
			log.Printf("channel error:%v", err)
			return
		}
		log.Printf("收到:%v\n", string(msg))
		chn.handlemsg(msgType, msg)

	}
}

//频道向客户端发送消息
func (chn *Channel) Send(msg []byte) {
	_ = chn.conn.WriteMessage(1, []byte(fmt.Sprintf("你的房间号:%v,你的ID:%v  ", chn.rm.name, chn.id)+string(msg)))
	log.Printf("发送:%v\n", string(msg))
	//考虑需不需要加锁，读和写冲不冲突。
}

//主动关闭频道
func (chn *Channel) Close() error {
	defer chn.rm.DelCh(chn.id)
	return chn.conn.Close()
}

//处理客户端发送的消息
func (chn *Channel) handlemsg(msgType int, msg []byte) {
	chn.Send(msg)
}

//将频道关联到房间
func (chn *Channel) LinkRoom(rm *Room) {
	chn.rm = rm
	if _, ok := rm.chn[chn.id]; !ok {
		rm.chn[chn.id] = chn
	}
}
