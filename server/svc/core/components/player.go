package components

import "time"

type Player struct {
	ID string
}

func NewPlayer(id string) Player {
	return Player{ID: id}
}

func (p *Player) Update(dt time.Duration) {}
