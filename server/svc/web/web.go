package web

import (
	"soikke.li/sol"
	"soikke.li/sol/log"

	"context"
	"fmt"
	"net/http"
)

type Config struct {
	sol.Component

	Port int
	TokenSecret string `yaml:"token_secret"`
	ClientPath string `yaml:"client_path"`

	db Datastore
}

func (cfg *Config) Init(log log.Logger) {
	cfg.Component.Init(`web`, log)

	if cfg.TokenSecret == `` {
		cfg.Log.Fatal().Msg(`missing token_secret, cannot start web service`)
	}
}

func (cfg *Config) InitDB(db Datastore) {
	cfg.db = db
}

func (cfg *Config) StartHub(ctx context.Context) *Hub {
	h := &Hub{
		clients:	make(map[*Client]bool),
		Incoming:	make(chan []byte),
		Broadcast:	make(chan []byte),
		Register:	make(chan *Client),
		Deregister:	make(chan *Client),
	}
	go h.run(ctx)
	return h
}

func (cfg *Config) StartHTTP(ctx context.Context, h *Hub) {
	root := http.NewServeMux()
	root.Handle(`/`, http.FileServer(http.Dir(`../client/dist`)))
	root.HandleFunc(`/users/`, cfg.UsersHandler)
	root.HandleFunc(`/users`, cfg.UsersHandler)
	root.HandleFunc(`/login`, cfg.LoginHandler)
	root.HandleFunc(`/ws`, func(w http.ResponseWriter, r *http.Request) {
		ServeWs(w, r, h)
	})
	go http.ListenAndServe(fmt.Sprintf(`:%d`, cfg.Port), root)
	// cfg.Log.Fatal().Err(http.ListenAndServe(fmt.Sprintf(`:%d`, cfg.Port), root))
}