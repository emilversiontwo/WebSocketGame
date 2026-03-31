package pages

import (
	"fmt"
	"log"
	"main/internal/transport/websocketServer"
	"net/http"

	_ "github.com/gorilla/websocket"
)

type Pages struct {
	webSocketServer *websocketServer.WebSocketServer
}

func NewPages(webSocketServer *websocketServer.WebSocketServer) *Pages {
	return &Pages{webSocketServer: webSocketServer}
}

func (p *Pages) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func (p *Pages) WsEndpoin(w http.ResponseWriter, r *http.Request) {
	p.webSocketServer.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := p.webSocketServer.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Client Connected")

	p.webSocketServer.Reader(ws)
}
