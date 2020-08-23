package main

import (
	"fmt"
	"os/exec"
	"os"
)

var buildTargets = `all sol`

func build(target string) (*exec.Cmd, error) {
	switch target {
	case `all`:
		return buildAll()
	case `sol`:
		return buildSol()
	}
	return nil, fmt.Errorf(`%s is not a valid build target. valid targets are: %s`, target, buildTargets)
}

func buildTarget(target string) (*exec.Cmd, error) {
	cmd := exec.Command(
		`go`,
		`build`,
		`-o`,
		`bin/`,
		target,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	return cmd, nil
}

func buildSol() (*exec.Cmd, error) {
	return buildTarget(`soikke.li/sol/cmd/sol`)
}

func buildAll() (*exec.Cmd, error) {
	return buildSol()
}