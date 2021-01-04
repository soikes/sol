package game

import (
	"soikke.li/sol/svc/core"
	"soikke.li/sol/svc/core/components"
)

func SpawnPlayerShip(id string, z *Zone) { //TODO look up ship associated with player from DB
	shipTrans := components.Transform{}
	shipPhys := components.NewPhysics(&shipTrans)
	shipPhys.AccelerationFactor = 1
	shipPhys.MaxSpeed = 10
	shipPhys.RotationFactor = 1
	shipInput := components.NewInput(&shipPhys)
	sp := components.NewPlayer(id)
	ship := core.NewEntity(&shipTrans, &shipPhys, &shipInput, &sp)
	z.Spawn(ship)
}
