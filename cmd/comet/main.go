package main

import (
	"geim/internal/comet"
	"geim/internal/comet/conf"
)

func main() {
	//comet 初始化
	instance := comet.InitComet(&conf.Comet{Room_size: 100})
	//rpc 服务
	go comet.InitRPCServer(&conf.Rpc{Port: 9001}, instance)
	go instance.InitRPCClient("127.0.0.1:9004")
	//websocket 服务
	comet.InitWebSocket(&conf.Websocket{Port: 9002}, instance)
	//退出

}
