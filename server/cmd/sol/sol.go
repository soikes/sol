package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"soikke.li/sol/config"

	"github.com/rs/zerolog/log"
)

var usage = `usage: sol command`

func main() {
	ctx := context.Background()
	ctx = handleSignals(ctx)
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
	<-ctx.Done()
	log.Info().Str(`cmd`, cmd).Msg(`shutting down`)
	if err != nil {
		log.Fatal().Err(err).Str(`cmd`, cmd).Msg(`failed to run`)
	}
}

func handleSignals(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	go func() {
		<-c
		cancel()
	}()
	return ctx
}