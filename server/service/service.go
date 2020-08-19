package service

import (
	"soikke.li/sol"
)

type Service struct {
	entities []sol.Entity
	systems  []sol.System
}

func Init() {
	
}

func (s *Service) Run() error {
	return nil
}