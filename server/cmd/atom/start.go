package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

var services = `crdb sol`

func start(ctx context.Context, svcs []string) error {
	ctx, cancel := context.WithCancel(ctx)
	var cmds []*exec.Cmd
	for _, svc := range svcs {
		fmt.Printf("attempting to start %s...\n", svc)
		switch svc {
		case `crdb`:
			cmd, err := startCRDB(ctx)
			if err != nil {
				cancel()
				return err
			}
			cmds = append(cmds, cmd)
		case `sol`:
			cmd, err := startSol(ctx)
			if err != nil {
				cancel()
				return err
			}
			cmds = append(cmds, cmd)
		default:
			return fmt.Errorf(`%s is not a valid service. valid services are: %s`, svc, services)
		}
	}
	return waitForCmds(ctx, cmds)
}

func waitForCmds(ctx context.Context, cmds []*exec.Cmd) error {
	ctx, cancel := context.WithCancel(ctx)
	var errs []error
	done := make(chan error)
	for _, cmd := range cmds {
		go func(cmd *exec.Cmd) {
			err := cmd.Wait()
			if err != nil {
				err = fmt.Errorf(`service %s failed: %w`, cmd.Path, err)
				cancel()
			}
			done<-err
		}(cmd)
	}
	for i := 0; i < len(cmds); i++ {
		err := <-done
		if err != nil {
			errs = append(errs, err)
		}
	}
	var err error
	if len(errs) > 0 {
		err = fmt.Errorf(`multiple errors: `)
		for _, e := range errs {
			err = errors.New(fmt.Sprint(err.Error(), e.Error()))
		}
	}
	return err
}

func startCRDB(ctx context.Context) (*exec.Cmd, error) {
	cmd := exec.CommandContext(
		ctx,
		`cockroach`,
		`start-single-node`,
		`--insecure`,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	return cmd, nil
}

func startSol(ctx context.Context, args ...string) (*exec.Cmd, error) {
	cmd := exec.CommandContext(
		ctx,
		`sol`,
		args...,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	return cmd, nil
}