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
