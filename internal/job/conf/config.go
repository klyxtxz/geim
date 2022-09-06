package conf

type Rpc struct {
	Port uint16
}
type RoomKafka struct {
	Addr  string
	Topic string
	Group string
	Name  string
}
type BroadCastKafka struct {
	Addr  string
	Topic string
}
