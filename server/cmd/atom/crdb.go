package main

import (
	"os/exec"
)

func startCRDB() (*exec.Cmd, error) {
	cmd := exec.Command(
		`cockroach`,
		`start-single-node`,
		`--insecure`,
	)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	return cmd, nil
}