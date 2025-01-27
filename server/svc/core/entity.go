package core

import (
	"time"

	"github.com/rs/xid"
)

type Entity struct {
	Id         string
	Components []Component
}

func NewEntity() Entity {
	return Entity{
		Id: xid.New().String(),
	}
}

func (e *Entity) Update(dt time.Duration) {
	for _, c := range e.Components {
		c.Update(dt)
	}
}

func (e *Entity) AddComponents(c ...Component) {
	for _, c := range c {
		c.Attach(e.Id)
		e.Components = append(e.Components, c)
	}
}
