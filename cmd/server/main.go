package main

import (
	"fmt"
	"log"
	"main/internal/game/game"
	"main/internal/transport/pages"
	"main/internal/transport/router"
	"main/internal/transport/websocketServer"
	"net/http"
)

func main() {
	fmt.Println("Hello WebSocket")
	Game := game.New()

	WebSocketServer := websocketServer.NewWebSocketServer(Game)

	go Game.Run()

	Pages := pages.NewPages(WebSocketServer)

	HttpRouter := router.NewRouter(WebSocketServer.Upgrader, Pages)
	HttpRouter.SetupRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
