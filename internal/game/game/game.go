package game

import (
	"log"
	"main/internal/game/message"
	"main/internal/game/player"
)

type Game struct {
	Players    map[string]*player.Player
	Register   chan *player.Player
	Unregister chan string
	Broadcast  chan message.Message
}

func New() *Game {
	return &Game{
		Players:    make(map[string]*player.Player),
		Register:   make(chan *player.Player),
		Unregister: make(chan string),
		Broadcast:  make(chan message.Message, 16),
	}
}

func (g *Game) Run() {
	for {
		select {
		case NewPlayer := <-g.Register:
			g.Players[NewPlayer.UUID] = NewPlayer
			log.Println("player joined game")
			// Сообщаем всем, что новый игрок присоединился
			msg := message.Message{
				Type: "player_joined",
				Payload: map[string]interface{}{
					"uuid": NewPlayer.UUID,
					"x":    NewPlayer.PosX,
					"y":    NewPlayer.PosY,
				},
			}

			g.Broadcast <- msg

		case OldPlayer := <-g.Unregister:
			if pl, ok := g.Players[OldPlayer]; ok {
				go delete(g.Players, pl.UUID)
				// Сообщаем всем, что новый игрок вышел
				msg := message.Message{
					Type: "player_destroyed",
					Payload: map[string]interface{}{
						"uuid": pl.UUID,
						"x":    pl.PosX,
						"y":    pl.PosY,
					},
				}
				g.Broadcast <- msg
			}

		case msg := <-g.Broadcast:
			for uuid, p := range g.Players {
				go func(uid string, pl *player.Player) {
					if err := pl.Conn.WriteJSON(msg); err != nil {
						log.Println("write error:", err)
						pl.Conn.Close()
						// помечаем на удаление
						g.Unregister <- pl.UUID
					}
				}(uuid, p)
			}
		}
	}
}
