package game

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
	"soikke.li/sol/message"
)

type Client struct {
	Id       string
	conn     *websocket.Conn
	hub      *Hub
	outgoing chan []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warn().Err(err).Msg(`failed to upgrade client ws connection`)
		http.Error(w, err.Error(), 400)
		return
	}
	z := GetZone()
	h := z.Hub //TODO this gets looked up for the current user. where are they? if they are new, they spawn in the home zone Sol
	cid := xid.New().String()
	c := Client{
		Id:       cid,
		conn:     conn,
		hub:      &h,
		outgoing: make(chan []byte),
	}
	re := message.Register{ID: cid}
	msg, err := re.Marshal()
	if err != nil {
		return
	}
	err = conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Warn().Err(err).Msg(`failed to send register message to client`)
		http.Error(w, err.Error(), 400)
		return
	}

	pid := xid.New().String()
	sid := SpawnPlayerShip(pid, z)
	sp := message.Spawn{ID: sid, Type: message.SpawnPlayer}

	msg, err = sp.Marshal()
	err = conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Warn().Err(err).Msg(`failed to send spawn message to client`)
		http.Error(w, err.Error(), 400)
		return
	}

	c.hub.RegisterClient(&c)

	go c.readPump()
	go c.writePump()
}

func (c *Client) readPump() {
	defer func() {
		c.hub.DeregisterClient(c)
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
		log.Info().Str(`msg`, string(message)).Msg(`got message`)
		c.hub.Incoming <- message
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(10 * time.Second)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case <-ticker.C:
			log.Info().Str(`id`, c.Id).Msg(`pinging client`)
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Error().Err(err).Msg(`failed to ping client`)
				c.hub.DeregisterClient(c)
				return
			}
		case msg := <-c.outgoing:
			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Error().Err(err).Msg(`failed to send msg to client`)
			}
		}
	}
}
