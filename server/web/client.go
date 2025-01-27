package web

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/rs/xid"
	"soikke.li/sol/msg"
)

type Client struct {
	id   xid.ID
	conn *websocket.Conn

	in  chan msg.Msg
	out chan msg.Msg
}

func NewClient(conn *websocket.Conn) Client {
	c := Client{
		id:   xid.New(),
		conn: conn,
		in:   make(chan msg.Msg),
		out:  make(chan msg.Msg),
	}
	return c
}

func (c *Client) Start() {
	go func() {
		for {
			out := <-c.out
			m, err := json.Marshal(out)
			if err != nil {
				log.Printf("Marshal: %v\n", err)
				continue
			}
			err = c.conn.WriteMessage(websocket.TextMessage, m)
			if err != nil {
				log.Printf("WriteMessage: %v\n", err)
				return
			}
		}
	}()

	// go func() {
	// 	for {
	// 		_, p, err := c.conn.ReadMessage()
	// 		if err != nil {
	// 			log.Printf("ReadMessage: %w\n", err)
	// 			return
	// 		}
	// 		c.in <-
	// 	}
	// }()
}
