package game

import (
	"soikke.li/sol/svc/core"
	"soikke.li/sol/svc/core/components"
)

func SpawnPlayerShip(pid string, z *Zone) string { //TODO look up ship associated with player from DB
	shipTrans := components.Transform{}
	shipPhys := components.NewPhysics(&shipTrans)
	shipPhys.AccelerationFactor = 8
	shipPhys.MaxSpeed = 15
	shipPhys.RotationFactor = 0.05
	shipInput := components.NewInput(&shipPhys)
	sp := components.NewPlayer(pid)
	ship := core.NewEntity()
	ship.AddComponents(&shipTrans, &shipPhys, &shipInput, &sp)
	z.Spawn(ship)

	return ship.Id
}
