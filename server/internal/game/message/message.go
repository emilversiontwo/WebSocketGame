package message

type Message struct {
	Type    string      `json:"type"`    // например: "join", "move", "leave", "chat"
	Payload interface{} `json:"payload"` // данные (могут быть разные структуры)
}

// Когда игрок двигается
type MovePayload struct {
	UUID string  `json:"uuid"`
	X    float32 `json:"x"`
	Y    float32 `json:"y"`
}

// Когда игрок присоединяется
type JoinPayload struct {
	UUID string `json:"uuid"`
}

// Когда игрок выходит
type LeavePayload struct {
	UUID string `json:"uuid"`
}

// Сообщение чата
type ChatPayload struct {
	From    string `json:"from"`
	Message string `json:"message"`
}
