package job

import (
	"geim/internal/job/conf"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

type RPC struct {
	ID string
}

func InitRPC(conf *conf.Rpc, object *RPC, done chan struct{}) {
	rpc.Register(object)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":"+strconv.Itoa(int(conf.Port)))
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	log.Print("RPC listening:", conf.Port)

	done <- struct{}{}
	http.Serve(l, nil)
	defer l.Close()
}

func (Rpc *RPC) FlushMap(arg map[string]map[string]struct{}, reply *string) error {
	log.Print("start FlushMap")
	for comet, _ := range arg {
		if _, ok := Cometrpc[comet]; !ok {
			go func(string) {
				client, err := rpc.DialHTTP("tcp", comet)
				log.Print("connect to comet: ", comet)
				if err != nil {
					log.Print("link to comet error :", err)
					return
				}
				Cometrpc[comet] = client
			}(comet)
		}
	}
	//是否需要深度复制?
	Comet2room = arg

	*reply = "OK"
	return nil
}
