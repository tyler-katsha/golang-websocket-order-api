package main

import (
	"log"
	"net/http"
	"tyler-katsha/golang-websocket-order-api/internals/handlers"
)

const (
	location = "http://localhost:7192"
	addr = ":7192"
)
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ws",handlers.WsHandler)
	mux.HandleFunc("/api/orders",handlers.CreateOrderHandler)

	fileServer := http.FileServer(http.Dir("./web"))
	mux.Handle("/",fileServer)


	log.Println("Server running on ",location)
	log.Fatal(http.ListenAndServe(addr,mux))

}