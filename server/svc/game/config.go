package game

import (
	"context"

	"soikke.li/sol"
	"soikke.li/sol/log"
)

type Config struct {
	sol.Component
}

func (cfg *Config) Init(l log.Logger) {
	cfg.Component.Init(`game`, l)
}

func (cfg *Config) Run(ctx context.Context) {
	cfg.Log.Info().Msg(`starting`)

	// Load persisted world here... many zones and components
	// z1 := GetZone() // TODO this comes from the DB

	// shipTrans := components.Transform{}
	// shipPhys := components.NewPhysics(&shipTrans)
	// shipPhys.AccelerationFactor = 1
	// shipPhys.MaxSpeed = 10
	// shipPhys.RotationFactor = 1
	// shipInput := components.NewInput(&shipPhys)
	// ship := core.NewEntity(&shipTrans, &shipPhys, &shipInput)
	// z1.Spawn(ship)
}
