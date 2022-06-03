package ws

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"time"
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

var WSChannel chan WSMessage

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
			keepAlive(cmd.Conn)
			log.Info("SW: Client Successfully Connected" + "[username: " + cmd.Username + "]")
		case TYPE_ALERT:
			log.Debug("cmd.Username = ", cmd.Username)
			addressee, exist := ws[cmd.Username]
			log.Debug("exist = ", exist)
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

func keepAlive(c *websocket.Conn) {
	//lastResponse := time.Now()
	//c.SetPongHandler(func(msg string) error {
	//	lastResponse = time.Now()
	//	return nil
	//})

	go func() {
		for {
			err := c.WriteMessage(websocket.PingMessage, []byte("keepalive"))
			if err != nil {
				return
			}
			time.Sleep(30 * 1000 * time.Millisecond)
			//if(time.Since(lastResponse) > timeout) {
			//	c.Close()
			//	return
			//}
		}
	}()
}

func NewWSServer() chan WSMessage {
	WSChannel = make(chan WSMessage, 0)
	go wsServer(WSChannel)
	return WSChannel
}
