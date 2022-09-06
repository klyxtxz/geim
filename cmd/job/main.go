package main

import (
	"geim/internal/job"
	"geim/internal/job/conf"
)

func main() {
	//启动本地服务,并上报master
	obj := job.InitJob(&conf.Rpc{Port: 9005})
	//连接kafka 开始推送
	go job.ConsumerGroup(&conf.RoomKafka{Name: obj.ID, Addr: "kafka-server:9092", Topic: "geim", Group: "test"})
	go job.KafkaGetBroadCast(&conf.BroadCastKafka{Addr: "kafka-server:9092", Topic: "broadcast"})
	select {}
	//退出

}
