package service

import (
	"soikke.li/sol/client"

	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Web struct {
	port int
}

func (w *Web) Run() {
	h := client.NewHub()
	go h.Run()
	http.Handle("/", http.FileServer(http.Dir("../client/dist")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		client.ServeWs(w, r, h)
	})
	log.Fatal().Err(http.ListenAndServe(fmt.Sprintf(`:%d`, w.port), nil))
}
