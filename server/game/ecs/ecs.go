package ecs

import (
	"time"

	"github.com/rs/xid"
)

type ECS struct {
	Entities map[string]Entity
}

type Index int

type Entity struct {
	id    xid.ID
	index Index
}

func NewEntity() Entity {
	return Entity{id: xid.New()}
}

func (e *Entity) String() string {
	return e.id.String()
}

func (e *Entity) Index() Index {
	return e.index
}

type TransformComponent struct {
	X float64
	Y float64
	Z float64
}

type VelocityComponent struct {
	VX float64
	VY float64
	VZ float64
}

type MovementComponent struct {
	Transform TransformComponent
	Velocity  VelocityComponent
}

type SystemManager struct {
	Systems []System
}

func (sm *SystemManager) Run() {
	last := time.Now()
	for {
		now := time.Now()
		dt := now.Sub(last)
		last = now
		for _, s := range sm.Systems {
			s.Compute(dt)
		}
	}
}

type System interface {
	Compute(dt time.Duration)
}

type MovementSystem struct {
	MovementComponents []MovementComponent
}

func (m *MovementSystem) Compute(dt time.Duration) {
	for i := 0; i < len(m.MovementComponents); i++ {
		mc := m.MovementComponents[i]
		mc.Velocity.VX = mc.Transform.X * float64(dt.Milliseconds())
		mc.Velocity.VY = mc.Transform.Y * float64(dt.Milliseconds())
		mc.Velocity.VZ = mc.Transform.Z * float64(dt.Milliseconds())
	}
}
