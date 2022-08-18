package main

import (
	"log"

	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

type Engine struct {
	Rms map[string]*Room
	mu  sync.Mutex
}

func (e *Engine) Initserver(port string) error {
	rms := e.Rms
	rms["first one"] = NewRoom(e, "first one")
	rm := rms["first one"]
	id := 1
	http.HandleFunc("/v1", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal("create websocket error:", err)
		}
		e.mu.Lock()
		chn := NewChannel(conn, strconv.Itoa(id))
		id++
		rm.AddCh(chn)
		e.mu.Unlock()
		go chn.Serve()
		rm.SendAll([]byte("广播"))

	})
	server := &http.Server{Addr: ":" + port, Handler: nil}
	server.ListenAndServe()
	defer server.Close()
	return nil
}
