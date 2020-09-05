package main

import (
	"context"
	"fmt"
	"os"

	"soikke.li/sol/config"

	"github.com/rs/zerolog/log"
)

var usage = `usage: sol command`

func main() {
	ctx := context.Background()
	// Register(`service`, cmdService)
	cfg := &config.Config{}
	err := cfg.Load(`etc/local.yml`)
	if err != nil {
		log.Fatal().Err(err).Msg(`failed to load configuration file`)
	}
	err = cfg.Init()
	if err != nil {
		log.Fatal().Err(err).Msg(`failed to initialize sol config`)
	}

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(usage)
		os.Exit(1)
	}
	cmd := args[0]
	// err := Call(cmd)
	switch cmd {
	case `service`:
		err = cmdService(ctx, cfg)
	case `setup`:
		err = cmdSetup(ctx, cfg)
	default:
		fmt.Printf("%s is not a valid command. exiting.\n", cmd)
		os.Exit(1)
	}
	if err != nil {
		fmt.Printf(fmt.Sprintf("failed to run %s\n", err.Error()))
		os.Exit(1)
	}
}