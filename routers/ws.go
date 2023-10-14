package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connectionPool = map[string]*websocket.Conn{}

type Message struct {
	MoveX   int    `json:"move_x"`
	MoveY   int    `json:"move_y"`
	Comment string `json:"comment"`
}

func initWsRouter(wr *gin.RouterGroup) {
	wr.GET("", func(c *gin.Context) {
		wsHandler(c.Writer, c.Request)
	})
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	id, _ := uuid.NewUUID()
	idStr := id.String()
	connectionPool[idStr] = conn
	defer delete(connectionPool, idStr)
	if err != nil {
		log.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}
	for {
		var mes *Message
		if err := conn.ReadJSON(&mes); err != nil {
			break
		}

		for _, c := range connectionPool {
			c.WriteJSON(mes)
		}
	}
}
