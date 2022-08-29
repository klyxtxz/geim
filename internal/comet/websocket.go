package comet

import (
	"geim/internal/comet/conf"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func (comet *Comet) httpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("收到连接!"))
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("create websocket error:", err)
		return
	}
	log.Print("连接ws成功")
	log.Print("cookie:", r.Header.Get("Cookie"))
	//鉴权

	room := new(Reply)
	err = comet.getroomid(r.Header.Get("Cookie"), room)
	if err != nil {
		log.Print("rpc:getroomid error:", err)
	}
	//鉴权成功,返回用户ID,房间号.
	log.Print(room, &conn)
	// comet.newchannel(conn, userId, roomId)

}
func InitWebSocket(conf *conf.Websocket, obj *Comet) {
	//start websocket server

	http.HandleFunc("/websocket", obj.httpHandler)
	server := &http.Server{Addr: ":" + strconv.Itoa(int(conf.Port)), Handler: nil}
	log.Print("WebSocket listening:", conf.Port)
	server.ListenAndServe()
	defer server.Close()
}

type Reply struct {
	Id   string
	Room string
}
