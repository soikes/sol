package client

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

type Client struct {
	Id        string
	conn      *websocket.Conn
	hub       *Hub
	outgoing  chan []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeWs(w http.ResponseWriter, r *http.Request, h *Hub) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warn().Err(err).Msg(`failed to upgrade client ws connection`)
		http.Error(w, err.Error(), 400)
		return
	}
	c := Client{
		Id: xid.New().String(),
		conn: conn,
		hub: h,
		outgoing: make(chan []byte),
	}
	err = conn.WriteMessage(websocket.TextMessage, []byte(c.Id))
	if err != nil {
		log.Warn().Err(err).Msg(`failed to send message`)
		http.Error(w, err.Error(), 400)
		return
	}
	c.hub.register <- &c

	go c.readPump()
	go c.writePump()
}

func (c *Client) readPump() {
	defer func() {
		c.hub.deregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Info().Err(err).Msg(`client disconnected unexpectedly`)
			}
			log.Info().Err(err).Msg(`client disconnected`)
			break
		}
		c.hub.incoming <- message
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(time.Second)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case <-ticker.C:
			log.Info().Str(`id`, c.Id).Msg(`pinging client`)
			c.conn.SetWriteDeadline(time.Now().Add(10*time.Second))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Error().Err(err).Msg(`failed to ping client`)
				return
			}
		}
	}
}