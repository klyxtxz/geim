package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime)
	e := Engine{Rms: make(map[string]*Room)}
	e.Initserver("3000")
}
