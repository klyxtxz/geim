package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime)
	e := Engine{Rms: make(map[string]*Room)}
	go InitRPC(&e)
	e.Initserver("3000")
}
