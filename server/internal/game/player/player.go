package player

import (
	"main/internal/game/message"

	"github.com/gorilla/websocket"
)

type Player struct {
	UUID string
	PosX float32
	PosY float32
	Conn *websocket.Conn
}

func (p *Player) Send(msg message.Message) error {
	return p.Conn.WriteMessage(1, []byte("you registered"))
}
