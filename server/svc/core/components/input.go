package components

import (
	"time"

	"soikke.li/sol/message"
)

type Input struct {
	Physics Physics

	pending message.Input
}

func (i *Input) Update(dt time.Duration) {
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
	i.pending = message.Input{} // TODO is this bad? How can we reset this value after update fires?
}

func (i *Input) QueueInput(m message.Input) {
	i.pending = m
}