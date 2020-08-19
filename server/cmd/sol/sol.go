package main

import (
	"fmt"
	"os"
)

var usage = `usage: sol command`

func main() {
	// Register(`service`, cmdService)

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(usage)
		os.Exit(1)
	}
	cmd := args[0]
	// err := Call(cmd)
	var err error
	switch cmd {
	case `service`:
		err = cmdService()
	default:
		fmt.Printf("%s is not a valid command. exiting.\n", cmd)
		os.Exit(1)
	}
	if err != nil {
		fmt.Printf(fmt.Sprintf("failed to run %s\n", err.Error()))
		os.Exit(1)
	}
}