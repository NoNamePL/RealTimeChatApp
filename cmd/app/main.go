package main

import (
	"fmt"
	"net/http"
	"realtimechatapp/pkg/client"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

func wsPage(res http.ResponseWriter, req *http.Request) {
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}).Upgrade(res, req, nil)

	if err != nil {
		http.NotFound(res, req)
		return
	}
	// client := &client.Client{id
	clientLocal := &client.Client{Id: uuid.NewV4().String(), Socket: conn, Send: make(chan []byte)}

	client.Manager.Register <- clientLocal

	go clientLocal.Read()
	go clientLocal.Write()

}

func main() {
	fmt.Println("Starting application...")
	go client.Manager.Start()
	http.HandleFunc("/ws", wsPage)
	http.ListenAndServe(":8080", nil)
}
