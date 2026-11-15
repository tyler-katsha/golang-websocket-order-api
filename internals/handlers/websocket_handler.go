package handlers

import (
	"context"
	"net/http"
	"log"
	"sync"
	"tyler-katsha/golang-websocket-order-api/internals"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

var (
	clients = make(map[*websocket.Conn]bool)
	clientsMu sync.Mutex
)
func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})

	if err != nil {
		log.Println("Upgrade error:",err)
		return
	}

	registerClient(conn)

	defer removeClient(conn)

	for {
		_,_,err := conn.Read(r.Context())

		if err != nil {
			log.Println("Client disconnected:", err)
			break
		}
	}

	
}

func registerClient(conn *websocket.Conn) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	clients[conn] = true
}

func removeClient(conn *websocket.Conn) {
	clientsMu.Lock()
	defer conn.CloseNow()
	delete(clients,conn)
	clientsMu.Unlock()
}

func BroadcastOrders(order internals.Order) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for conn := range clients {
		go pushOrders(conn,order)
	}
}

func pushOrders(conn *websocket.Conn, order internals.Order) {
	if err := wsjson.Write(context.Background(), conn, order); err != nil {
		log.Println("Write error: ", err)
	}
}