package comet

import (
	"geim/internal/comet/conf"
	"log"
	"net/rpc"
)

func (comet *Comet) InitLogic(conf conf.LogicAddr) error {
	client, err := rpc.DialHTTP("tcp", string(conf))
	comet.logicserver = client
	log.Print("InitLogic...")
	if err != nil {
		log.Fatal("link to logic error :", err)
		return err
	}
	return nil
}

func (comet *Comet) InitMaster(conf conf.MasterAddr) error {
	client, err := rpc.DialHTTP("tcp", string(conf))
	comet.masterserver = client
	log.Print("InitMaster...")
	if err != nil {
		log.Fatal("link to logic error :", err)
		return err
	}
	return nil
}

func (comet *Comet) getroomid(id string, reply *Reply) error {

	err := comet.logicserver.Call("Logic.CheckToken", id, reply)
	log.Print(reply)
	if err != nil {

		return err
	}
	return nil
}
