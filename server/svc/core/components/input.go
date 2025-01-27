package components

import (
	"time"

	"github.com/rs/zerolog/log"
	"soikke.li/sol/message"
)

type Input struct {
	Physics *Physics

	entity    string
	isPending bool
	pending   message.Input
}

func (i *Input) Attach(entity string) {
	i.entity = entity
}

func NewInput(p *Physics) Input {
	return Input{Physics: p}
}

func (i *Input) Update(dt time.Duration) {
	if i.isPending {
		log.Info().Interface(`pending`, i.pending).Msg(`update`)
	}
	if i.pending.ForwardPress {
		i.Physics.Accelerate()
	} else {
		i.Physics.StopAccelerating()
	}

	if i.pending.LeftPress {
		i.Physics.RotateY(1)
	} else if i.pending.RightPress {
		i.Physics.RotateY(-1)
	} else {
		i.Physics.StopRotating()
	}
	i.isPending = false
	i.pending = message.Input{} // TODO is this bad? How can we reset this value after update fires?
}

func (i *Input) QueueInput(m message.Input) {
	i.isPending = true
	i.pending = m //TODO make this a queue that drops messages that are too old - prevent late input from being processed, makes driving suck
}
