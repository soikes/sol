package game

import (
	"soikke.li/sol"
)

type Config struct {
	entities []sol.Entity
	systems  []sol.System
}

func Init() {
	
}

func (c *Config) Run() error {
	return nil
}