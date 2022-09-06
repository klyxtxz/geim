package main

import (
	"geim/internal/master"
	"geim/internal/master/conf"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	obj := master.InitMaster()
	//rpc 服务
	//接受comet心跳
	go master.InitRPC(&conf.Rpc{Port: 9006}, obj)
	//接受job请求
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-c
	//退出

}
