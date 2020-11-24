package core

import (
	"time"

	"github.com/rs/xid"
)

type Entity struct {
	Id string
	Components []Component
}

func NewEntity(components ...Component) Entity {
	return Entity{
		Id: xid.New().String(),
		Components: components,
	}
}

func (e *Entity) Update(dt time.Duration) {
	for _, c := range e.Components {
		c.Update(dt)
	}
}

func (e *Entity) AddComponent(c Component) {
	e.Components = append(e.Components, c)
}