package conf

type HTTPConf struct {
	Port uint16
}
type Rpc struct {
	Port uint16
}
type LogicConf struct {
	Kafka string
	Redis string
}
type RoomKafka struct {
	Addr  string
	Topic string
	Group string
}
type BroadKafka struct {
	Addr  string
	Topic string
}
