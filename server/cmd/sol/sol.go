package main

import (
	"context"
	"fmt"
	"os"

	"soikke.li/sol"
)

var usage = `usage: sol command`

func main() {
	ctx := context.Background()
	// Register(`service`, cmdService)
	cfg := &sol.Config{}
	err := cfg.Load(``)
	if err != nil {
		panic(err)
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