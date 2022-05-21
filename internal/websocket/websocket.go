package ws

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

const (
	TYPE_NEW_CONNECTION = 0
	TYPE_ALERT          = 1

	STATUS_INFO    = 0
	STATUS_WARNING = 1
)

type WSMessage struct {
	Type          int
	Username      string
	Conn          *websocket.Conn
	Message       string
	MessageStatus int
}

func wsServer(in chan WSMessage) {
	log.Info("WS: server start")
	ws := map[string]*websocket.Conn{}
	for {
		cmd, opened := <-in
		if !opened {
			log.Info("WS: server stop")
			return
		}

		switch cmd.Type {
		case TYPE_NEW_CONNECTION:
			ws[cmd.Username] = cmd.Conn
			log.Info("SW: Client Successfully Connected" + "[username: " + cmd.Username + "]")
		case TYPE_ALERT:
			addressee, exist := ws[cmd.Username]
			if !exist {
				continue
			}
			if err := addressee.WriteMessage(1, []byte(cmd.Message)); err != nil {
				log.Info(err)
				continue
			}
			log.Info("WS: send message")
		}
	}
}

func NewWSServer() chan WSMessage {
	in := make(chan WSMessage, 0)
	go wsServer(in)
	return in
}
