package channel

import (
	"log"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type Channel struct {
	Conn *websocket.Conn
	Id   string
}

func (chn *Channel) Serve() {
	chn.Conn.SetCloseHandler(chn.close)
	for {
		time.Sleep(time.Second)
		_, msg, err := chn.Conn.ReadMessage()
		if err != nil {
			if strings.Contains(err.Error(), "close") {
				log.Printf("连接关闭")

				return
			}
			log.Printf("channel error:%v", err)
			return
		}
		_ = chn.Conn.WriteMessage(1, msg)
		log.Printf("回发:%v\n", string(msg))
	}
}

func (chn *Channel) Send(msg string) {
	_ = chn.Conn.WriteMessage(1, []byte(msg))
	//考虑需不需要加锁，读和写冲不冲突。
}

func (chn *Channel) close(code int, msg string) error {

	return nil
}

// func (chn *Channel) Serve() {
// 	for msg := range chn.Tunnel {
// 		chn.Conn.WriteMessage(1, []byte(msg))
// 		log.Printf("sended:%v\n", msg)
// 	}
// }
