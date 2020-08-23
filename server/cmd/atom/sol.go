package main

import (
	"os/exec"
)

func startSol(args ...string) (*exec.Cmd, error) {
	cmd := exec.Command(
		`sol`,
		args...,
	)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	return cmd, nil
}