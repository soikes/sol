package web

import (
	"soikke.li/sol/svc/client"

	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Config struct {
	Port int
	TokenSecret string `yaml:"token_secret"`

	db Datastore
}

func (cfg *Config) InitDB(db Datastore) {
	cfg.db = db
}

func (cfg *Config) Run(ctx context.Context) {
	h := client.NewHub()
	go h.Run(ctx)
	http.Handle("/", http.FileServer(http.Dir("../client/dist")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		client.ServeWs(w, r, h)
	})
	http.HandleFunc("/users", cfg.UsersHandler)
	log.Fatal().Err(http.ListenAndServe(fmt.Sprintf(`:%d`, cfg.Port), nil))
}
