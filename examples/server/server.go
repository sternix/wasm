package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ws_time(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	c := time.Tick(1 * time.Second)
	for now := range c {
		writer, err := conn.NextWriter(websocket.TextMessage)
		if err != nil {
			log.Println(err)
			return
		}
		writer.Write([]byte(now.Format("15:04:05")))
		writer.Close()
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path[1:]) > 0 {
		http.ServeFile(w, r, r.URL.Path[1:])
	} else {
		http.ServeFile(w, r, "assets/index.html")
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ws_time", ws_time)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
