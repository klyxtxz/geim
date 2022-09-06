package job

import (
	"encoding/json"
	"log"
)

type httpreq struct {
	CallerId   string
	CallerRoom string
	Operation  string
	Target     string
	Content    string
}

func handleRoomMsg(msg []byte) {
	req := new(httpreq)
	err := json.Unmarshal(msg, req)
	if err != nil {
		log.Print("unmarshal error:", err)
		return
	}
	switch req.Operation {
	case "1":
		for comet, rm := range Comet2room {
			if _, ok := rm[req.Target]; ok {
				pushRoom(comet, req.Target, req.Content)
			}
		}
	default:
		return
	}
}
func handleBroadCast(msg []byte) {
	req := new(httpreq)
	err := json.Unmarshal(msg, req)
	if err != nil {
		log.Print("unmarshal error:", err)
		return
	}
	log.Print(Comet2room)
	for comet, _ := range Comet2room {
		pushBroadCast(comet, req.Content)
	}

}

type Pushroom struct {
	Target  string
	Content string
}

func pushRoom(comet, target, content string) {
	room := Pushroom{Target: target, Content: content}
	var reply string
	err := Cometrpc[comet].Call("Comet.PushRoom", room, &reply)
	if err != nil {
		log.Printf("pushRoom err: %v\n", err)
	}
	log.Print("pushRoom success:", reply)
}
func pushBroadCast(comet, content string) {
	var reply string
	err := Cometrpc[comet].Call("Comet.PushAll", content, &reply)
	if err != nil {
		log.Printf("pushBroadCast err: %v\n", err)
	}
	log.Print("pushBroadCast success", reply)
}
