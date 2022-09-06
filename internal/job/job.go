package job

import (
	"geim/internal/job/conf"
	"log"
	"net/rpc"
	"os"
)

var Comet2room map[string]map[string]struct{}
var Cometrpc map[string]*rpc.Client

func InitJob(conf *conf.Rpc) *RPC {

	self, ok := os.LookupEnv("SELF")
	if !ok {
		log.Fatal("get self address error")
	}
	master, ok := os.LookupEnv("MASTER")
	if !ok {
		log.Fatal("get master address error")
	}

	Comet2room = make(map[string]map[string]struct{})
	Cometrpc = make(map[string]*rpc.Client)
	rpc := new(RPC)
	rpc.ID = self
	// 连接master 获取comet-room
	done := make(chan struct{})
	go InitRPC(conf, rpc, done)
	<-done
	callmaster(master, self)
	//向ｍａｓｔｅｒ上报上线消息

	return rpc
}
func callmaster(master, self string) {
	client, err := rpc.DialHTTP("tcp", master)
	log.Print("connect to master:", master)
	if err != nil {
		log.Fatal("link to master error :", err)
		return
	}
	reply := ""
	err = client.Call("Master.Job", self, &reply)
	if err != nil {
		log.Fatal("push master error :", err)
		return
	}
	log.Print("master reply:", reply)
	defer client.Close()
}
