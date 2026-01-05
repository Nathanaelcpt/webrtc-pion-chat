package signaling

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWS(w http.ResponseWriter, r *http.Request, handler func([]byte)) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		handler(msg)
	}
}
