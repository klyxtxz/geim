package main

import (
	"geim/internal/logic"
	"geim/internal/logic/conf"
)

func main() {
	//logic 服务
	lg := logic.InitLogic(conf.LogicConf{Kafka: "kafka-server:9092", Redis: "redis-server:6379"})
	//rpc 服务
	go logic.InitRPC(&conf.Rpc{Port: 9004}, lg)
	//http 服务
	logic.InitHTTPServer(&conf.HTTPConf{Port: 9003}, lg)
	//退出
	defer lg.Close()
}
