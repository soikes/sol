package web

import (
	"soikke.li/sol"
	"soikke.li/sol/log"
	"soikke.li/sol/svc/client"

	"context"
	"fmt"
	"net/http"
)

type Config struct {
	sol.Component

	Port int
	TokenSecret string `yaml:"token_secret"`

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

func (cfg *Config) Run(ctx context.Context) {
	h := client.NewHub()
	go h.Run(ctx)

	root := http.NewServeMux()
	root.Handle(`/`, http.FileServer(http.Dir(`../client/dist`)))
	root.HandleFunc(`/users/`, cfg.UsersHandler)
	root.HandleFunc(`/users`, cfg.UsersHandler)
	root.HandleFunc(`/login`, cfg.LoginHandler)
	root.HandleFunc(`/ws`, func(w http.ResponseWriter, r *http.Request) {
		client.ServeWs(w, r, h)
	})
	cfg.Log.Fatal().Err(http.ListenAndServe(fmt.Sprintf(`:%d`, cfg.Port), root))
}