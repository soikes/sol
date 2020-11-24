package web

import (
	"context"
)

type Hub struct {
	Incoming	chan []byte
	Broadcast	chan []byte
	Register	chan *Client
	Deregister	chan *Client

	clients		map[*Client]bool
}

func (h *Hub) run(ctx context.Context) error {
	for {
		select {
		case c := <-h.Register:
			h.clients[c] = true;
		case c := <-h.Deregister:
			if _, ok := h.clients[c]; ok {
				delete(h.clients, c)
				close(c.outgoing)
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}