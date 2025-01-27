package components

import "time"

type Player struct {
	ID string

	entity string
}

func (p *Player) Attach(entity string) {
	p.entity = entity
}

func NewPlayer(id string) Player {
	return Player{ID: id}
}

func (p *Player) Update(dt time.Duration) {}
