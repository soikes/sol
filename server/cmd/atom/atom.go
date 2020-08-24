package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var usage = `usage: atom command [-- command2]`
var commands = `build start`

func main() {
	ctx := context.Background()
	handleSignals(ctx)

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(usage)
		os.Exit(1)
	}
	cmd := args[0]
	// err := Call(cmd)
	// var err error
	switch cmd {
	case `build`:
		if len(args) < 2 {
			fmt.Println(`build: builds targets. usage: atom build target [target2...]`)
			os.Exit(1)
		}
		target := args[1]
		cmd, err := build(target)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = cmd.Wait()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	case `start`:
		if len(args) < 2 {
			fmt.Println(`start: starts services in order. usage: atom start service [service2...]`)
			os.Exit(1)
		}
		svcs := args[1:]
		err := start(ctx, svcs)
		if err != nil {
			fmt.Println(`start failed: `, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Printf("%s is not a valid command. valid commands are: %s\n", cmd, commands)
		os.Exit(1)
	}
	// if err != nil {
	// 	fmt.Printf(fmt.Sprintf("failed to run %s\n", err.Error()))
	// 	os.Exit(1)
	// }
}

func handleSignals(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	go func() {
		<-c
		cancel()
	}()
}