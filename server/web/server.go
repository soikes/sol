package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"soikke.li/sol/msg"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Server struct {
	clients  map[string]*Client
	register chan *Client
}

func (s *Server) Start() error {
	s.register = make(chan *Client)

	go s.pumpClientRegister()
	return s.startHTTP()
}

func (s *Server) startHTTP() error {
	http.HandleFunc(`/ws`, func(w http.ResponseWriter, r *http.Request) {
		serveWs(s.register, w, r)
	})
	addr := "0.0.0.0:9000"
	srv := http.Server{Addr: addr}
	log.Printf("listening on %s\n", addr)
	return srv.ListenAndServe()
}

func (s *Server) pumpClientRegister() {
	s.clients = make(map[string]*Client)
	for {
		select {
		case c := <-s.register: //TODO handle deregister as well. Need to check msg type
			id := c.id.String()
			log.Printf("new client %s\n", id)
			s.clients[id] = c
			rm := msg.RegisterMsg{ClientID: id}
			b, err := json.Marshal(rm)
			if err != nil {
				log.Printf("Marshal: %v", err)
				continue
			}
			m := msg.Msg{Type: msg.RegisterMsgType, Data: b}
			c.out <- m
		}
	}
}

func serveWs(register chan *Client, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade: %v", err)
		return
	}
	c := NewClient(conn)
	c.Start()
	register <- &c
}
