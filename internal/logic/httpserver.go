package logic

import (
	"fmt"
	"geim/internal/logic/conf"
	"log"
	"net/http"
	"strconv"
)

type httpreq struct {
	CallerId   string
	CallerRoom string
	Operation  string
	Target     string
	Content    string
}

func InitHTTPServer(conf *conf.HTTPConf, obj *Logic) {
	http.HandleFunc("/", obj.HTTPHandler)
	log.Print("HTTP listening:", conf.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(int(conf.Port)), nil))
}

func (logic *Logic) HTTPHandler(w http.ResponseWriter, req *http.Request) {
	if len(req.Header["Id"])*len(req.Header["Rm"])*len(req.Header["Op"])*len(req.Header["Tg"])*len(req.Header["Ct"]) == 0 {
		log.Print("http request error: header error")
		return
	}
	msg := httpreq{
		CallerId:   req.Header["Id"][0],
		CallerRoom: req.Header["Rm"][0],
		Operation:  req.Header["Op"][0],
		Target:     req.Header["Tg"][0],
		Content:    req.Header["Ct"][0],
	}
	log.Print(msg)

	//redis put id room
	err := logic.putredis(msg.CallerId, msg.CallerRoom)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	//kafka push msg
	if msg.Operation == "0" {

		err = logic.pushkafka(msg, "broadcast")
	} else {
		err = logic.pushkafka(msg, "geim")
	}
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
}
