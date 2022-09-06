package conf

type Websocket struct {
	Port uint16
}

type Rpc struct {
	Port uint16
}

type LogicAddr string
type MasterAddr string

type Comet struct {
	Comet_id     string
	Room_size    uint16
	Channel_size uint16
}
