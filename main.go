package main

import (
	"log"
	"main/room"
	socket "main/websocket"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime)
	e := socket.Engine{Rms: make(map[string]*room.Room)}
	e.Initserver("3000")
}
