package websocketServer

import (
	"encoding/json"
	"log"
	"main/internal/game/game"
	"main/internal/game/message"
	"main/internal/game/player"

	"github.com/gorilla/websocket"
)

type WebSocketServer struct {
	game     *game.Game
	Upgrader *websocket.Upgrader
}

func NewWebSocketServer(game *game.Game) *WebSocketServer {
	return &WebSocketServer{
		game: game,
		Upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (ws *WebSocketServer) Reader(conn *websocket.Conn) {
	for {
		var msg message.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			return
		}

		switch msg.Type {
		case "move":
			var payload message.MovePayload
			DecodePayload(msg.Payload, &payload)
			// обновляем координаты
			if _, ok := ws.game.Players[payload.UUID]; ok {
				ws.game.Players[payload.UUID].PosX += payload.X
				ws.game.Players[payload.UUID].PosY += payload.Y
				// Сообщаем всем, что игрок двигрался
				msg := message.Message{
					Type: "player_moved",
					Payload: map[string]interface{}{
						"uuid": ws.game.Players[payload.UUID].UUID,
						"x":    ws.game.Players[payload.UUID].PosX,
						"y":    ws.game.Players[payload.UUID].PosY,
					},
				}
				ws.game.Broadcast <- msg
			}
			// отправляем другим игрокам
		case "chat":
			var payload message.ChatPayload
			DecodePayload(msg.Payload, &payload)
			ws.game.Broadcast <- message.Message{
				Type:    "chat",
				Payload: payload,
			}
		case "leave":
			var payload message.LeavePayload
			DecodePayload(msg.Payload, &payload)
			ws.game.Unregister <- payload.UUID
		case "join":
			var payload message.JoinPayload
			DecodePayload(msg.Payload, &payload)
			NewPlayer := &player.Player{
				UUID: payload.UUID,
				PosX: 0,
				PosY: 0,
				Conn: conn,
			}
			log.Println("player joined game in ws")

			ws.game.Register <- NewPlayer
		default:
			log.Printf("unknown message type: %s", msg.Type)
			conn.WriteMessage(1, []byte("error"))
		}
	}
	log.Printf("error cicle end")
}

func DecodePayload(payload interface{}, out interface{}) error {
	bytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, out)
}
