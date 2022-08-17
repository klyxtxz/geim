package room

import (
	"main/channel"
)

type Room struct {
	name string
	Chn  map[string]*channel.Channel
}

func (rm *Room) SendAll(msg string) {
	for _, chn := range rm.Chn {
		chn.Send(msg + "你的ID是:" + chn.Id)
	}
}
