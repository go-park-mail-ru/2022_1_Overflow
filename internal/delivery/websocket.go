package delivery

import (
	"OverflowBackend/internal/session"
	ws "OverflowBackend/internal/websocket"
	"OverflowBackend/pkg"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (d *Delivery) WSConnect(w http.ResponseWriter, r *http.Request) {
	log.Info("WebSocket Endpoint")

	data, err := session.Manager.GetData(r)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.SESSION_ERR)
		return
	}

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warning(err)
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}

	if err := conn.WriteMessage(1, []byte("okey")); err != nil {
		log.Warning(err)
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}

	ws.WSChannel <- ws.WSMessage{
		Type:     ws.TYPE_NEW_CONNECTION,
		Username: data.Username,
		Conn:     conn,
	}
}
