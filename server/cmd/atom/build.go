package main

import (
	"fmt"
	"os"
	"os/exec"
)

var buildTargets = `all sol`

func build(target string) (*exec.Cmd, error) { //TODO plumb contexts
	switch target {
	case `all`:
		return buildAll()
	case `sol`:
		return buildSol()
	case `client`:
		return buildClient()
	}
	return nil, fmt.Errorf(`%s is not a valid build target. valid targets are: %s`, target, buildTargets)
}

func buildGo(target string) (*exec.Cmd, error) {
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

func buildNPM(target string) (*exec.Cmd, error) {
	cmd := exec.Command(
		`npm`,
		`run`,
		`build`,
		`--prefix`,
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
	return buildGo(`soikke.li/sol/cmd/sol`)
}

func buildClient() (*exec.Cmd, error) {
	return buildNPM(`../client`)
}

//TODO return slice of cmds
func buildAll() (*exec.Cmd, error) {
	c, err := buildClient()
	if err != nil {
		return nil, err
	}
	c.Wait()
	return buildSol()
}
