package master

import (
	"log"
	"net/rpc"
)

//处理job上线

func (mst *Master) Job(arg string, rep *string) error {

	mst.AddJobrpc(arg)

	go mst.flushmapone(arg)
	*rep = "OK"
	return nil
}

func iselementexist(e string, list []string) bool {
	for _, i := range list {
		if i == e {
			return true
		}
	}
	return false
}

func (mst *Master) AddJobrpc(job string) {
	client, err := rpc.DialHTTP("tcp", job)
	log.Print("connect to job:", job)
	if err != nil {
		log.Print("link to job error :", err)
		return
	}
	mst.Joblist[job] = client
}
