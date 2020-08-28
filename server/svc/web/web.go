package web

import (
	"soikke.li/sol/client"

	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Config struct {
	Port int `yaml:port`
}

func (c *Config) Init() error { return nil }

func (c *Config) Run(ctx context.Context) {
	h := client.NewHub()
	go h.Run(ctx)
	http.Handle("/", http.FileServer(http.Dir("../client/dist")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		client.ServeWs(w, r, h)
	})
	log.Fatal().Err(http.ListenAndServe(fmt.Sprintf(`:%d`, c.Port), nil))
}
