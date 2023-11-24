package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		conn, err := upgrader.Upgrade(w, r, nil)

		checkErr(err)
		for {
			msgType, msg, err := conn.ReadMessage()

			checkErr(err)

			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
			if err = conn.WriteMessage(msgType, []byte(fmt.Sprintf("%s: %s", conn.RemoteAddr(), msg))); err != nil {
				log.Fatal(err)
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	} else {
		return
	}
}
