package game

import (
	"context"
)

type Hub struct {
	Incoming   chan []byte
	Broadcast  chan []byte
	register   chan *Client
	deregister chan *Client

	clients map[*Client]bool //TODO use slice for performance
}

func NewHub() Hub {
	h := Hub{
		clients:    make(map[*Client]bool),
		Incoming:   make(chan []byte),
		Broadcast:  make(chan []byte),
		register:   make(chan *Client),
		deregister: make(chan *Client),
	}
	go h.Start(context.Background())
	return h
}

func (h *Hub) Start(ctx context.Context) error {
	for {
		select {
		case c := <-h.register:
			h.clients[c] = true
		case c := <-h.deregister:
			if _, ok := h.clients[c]; ok {
				delete(h.clients, c)
				close(c.outgoing)
			}
		case msg := <-h.Broadcast:
			for c := range h.clients {
				c.outgoing <- msg
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (h *Hub) RegisterClient(c *Client) {
	h.register <- c
}

func (h *Hub) DeregisterClient(c *Client) {
	h.deregister <- c
}
