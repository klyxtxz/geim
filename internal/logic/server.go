package logic

import "geim/internal/logic/conf"

type Logic struct {
	redis *Redis
	kafka *kafka
}

func (logic *Logic) putredis(id, room string) error {
	return logic.redis.put(id, id, room)
}
func (logic *Logic) pushkafka(req httpreq, topic string) error {
	return logic.kafka.push(req, topic)

}

//"redis-server:6379"
//"kafka-server:9092"
func InitLogic(conf conf.LogicConf) (logic *Logic) {
	logic = new(Logic)
	logic.kafka = Initkafka(conf.Kafka)
	logic.redis = Initredis(conf.Redis)
	return logic
}
func (logic *Logic) Close() {
	logic.kafka.close()
	logic.redis.rdb.Close()
}
