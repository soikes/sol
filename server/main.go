package main

import (
	"net/http"
	//"math/rand"

	"github.com/rs/zerolog/log"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("../client/dist")))
	// http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Print(`connected to ws`)
	// 	server.ServeWs(universe, hub, w, r)
	// })
	log.Fatal().Err(http.ListenAndServe(":8080", nil))
}
