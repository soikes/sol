package main

import(
	"fmt"

	"soikke.li/sol/service"
)

func cmdService() error {
	fmt.Println(`running service`)
	svc := service.Service{}
	svc.Run()
	return nil
}