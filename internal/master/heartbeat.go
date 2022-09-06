package master

import "log"

type Comet struct {
	Addr  string
	Rooms []string
}

func (mst *Master) HeartBeat(arg *Comet, rep *string) error {
	go mst.checkcomet(*arg)
	*rep = "OK"
	log.Print("get heartbeat:", arg.Addr, arg.Rooms)
	return nil
}
func (mst *Master) checkcomet(arg Comet) {
	comet := make(map[string]struct{})
	for _, rm := range arg.Rooms {
		comet[rm] = struct{}{}
	}
	if len(mst.CometRoom[arg.Addr]) == len(arg.Rooms) {
		for _, i := range arg.Rooms {
			if _, ok := mst.CometRoom[arg.Addr][i]; !ok {
				mst.CometRoom[arg.Addr] = comet
				mst.flushmapall()
				log.Print("房间列表变化,刷新所有job")

				return
			}

		}
		log.Print("房间列表没有变化")
		return
	}
	mst.CometRoom[arg.Addr] = comet
	mst.flushmapall()
	log.Print("房间列表变化,刷新所有job")

}
