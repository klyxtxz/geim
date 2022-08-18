package main

type Room struct {
	name string
	chn  map[string]*Channel
	en   *Engine
}

// 新建一个Room 对象
func NewRoom(en *Engine, name string) *Room {
	return &Room{name: name, chn: map[string]*Channel{}, en: en}
}

//给Room 对象添加频道
func (rm *Room) AddCh(ch *Channel) {
	rm.chn[ch.id] = ch
	ch.LinkRoom(rm)
}

//给Room 对象删除频道
func (rm *Room) DelCh(id string) {
	delete(rm.chn, id)
}

//向Room 对象所有频道发送消息
func (rm *Room) SendAll(msg []byte) {
	for _, chn := range rm.chn {
		chn.Send(msg)
	}
}

//给Room 对象中指定ID对象发送消息
func (rm *Room) SendOne(id string, msg []byte) {
	rm.chn[id].Send(msg)
}

//主动关闭房间
func (rm *Room) Close() {
	delete(rm.en.Rms, rm.name)
}
