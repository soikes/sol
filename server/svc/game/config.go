package game

import (
	"context"

	"soikke.li/sol"
	"soikke.li/sol/log"
	"soikke.li/sol/svc/core"
	"soikke.li/sol/svc/core/components"
	"soikke.li/sol/svc/web"
	"soikke.li/sol/svc/zone"
)

type Config struct {
	sol.Component
}

func (cfg *Config) Init(l log.Logger) {
	cfg.Component.Init(`game`, l)
}

func (cfg *Config) Run(ctx context.Context, h *web.Hub) {
	cfg.Log.Info().Msg(`starting`)
	z1 := zone.Zone{}
	z1.Init(cfg.Component.Log)
	z1.SubscribeIncoming(h.Incoming)

	shipTrans := components.Transform{}
	shipPhys := components.NewPhysics(&shipTrans)
	ship := core.NewEntity(&shipTrans, &shipPhys)
	z1.Spawn(ship)
}