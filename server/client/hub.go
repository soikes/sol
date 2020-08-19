package client

type Hub struct {
	clients		map[*Client]bool
	incoming	chan []byte
	broadcast	chan []byte
	register	chan *Client
	deregister	chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:	make(map[*Client]bool),
		incoming:	make(chan []byte),
		broadcast:	make(chan []byte),
		register:	make(chan *Client),
		deregister:	make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true;
		case client := <-h.deregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.outgoing)
			}
		}
	}
}