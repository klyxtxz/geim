package master

import (
	"log"
	"net/rpc"
)

type Master struct {
	CometRoom map[string]map[string]struct{}
	Joblist   map[string]*rpc.Client
}

func InitMaster() *Master {
	mst := new(Master)
	mst.CometRoom = map[string]map[string]struct{}{}
	mst.Joblist = make(map[string]*rpc.Client)
	log.Print("initing master")
	return mst
}

func (mst *Master) flushmapone(job string) {
	var reply string
	mst.Joblist[job].Call("RPC.FlushMap", mst.CometRoom, &reply)
	log.Print("flushmapone:", reply)
}
func (mst *Master) flushmapall() {
	var reply string
	for _, client := range mst.Joblist {
		client.Call("RPC.FlushMap", mst.CometRoom, &reply)
	}
	log.Print("complete flushall")
}
