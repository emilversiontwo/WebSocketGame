package router

import (
	"main/internal/transport/pages"
	"net/http"

	"github.com/gorilla/websocket"
)

type Router struct {
	pages    *pages.Pages
	upgrader *websocket.Upgrader
}

func NewRouter(upgrader *websocket.Upgrader, pages *pages.Pages) *Router {
	return &Router{
		pages:    pages,
		upgrader: upgrader,
	}
}

func (r *Router) SetupRoutes() {
	http.HandleFunc("/", r.pages.HomePage)
	http.HandleFunc("/ws", r.pages.WsEndpoin)
}
