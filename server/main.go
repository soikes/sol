package main

import (
	"log"

	"soikke.li/sol/web"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	server := web.Server{}
	log.Println(server.Start())
}
