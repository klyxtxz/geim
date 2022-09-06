package main

import (
	"geim/internal/comet"
	"geim/internal/comet/conf"
	"log"
	"time"
)

func main() {
	//comet 初始化
	instance := comet.InitComet(&conf.Comet{Room_size: 100, Comet_id: "127.0.0.1:9001"})
	//rpc 服务
	go comet.InitRPCServer(&conf.Rpc{Port: 9001}, instance)
	go instance.InitLogic("127.0.0.1:9004")
	//websocket 服务
	go comet.InitWebSocket(&conf.Websocket{Port: 9002}, instance)

	//每5秒心跳
	timer := time.NewTicker(5 * time.Second)
	instance.InitMaster("127.0.0.1:9006")
	for {
		<-timer.C
		log.Print("timer tick")
		instance.Heartbeat()
	}
	//退出

}
