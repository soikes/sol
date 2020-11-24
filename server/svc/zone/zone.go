package zone

import (
	"soikke.li/sol"
	"soikke.li/sol/log"
	"soikke.li/sol/message"
	"soikke.li/sol/svc/core"
	"soikke.li/sol/svc/core/loop"
)

type Zone struct {
	sol.Component

	Loop loop.Config

	// Outgoing chan message.Message
	// Incoming chan message.Input
}

func (z *Zone) Init(log log.Logger) {
	z.Component.Init(`zone`, log)

	z.Loop = loop.Config{}
	z.Loop.Init(log)

	// z.Incoming = make(chan message.Input)
	// z.Outgoing = make(chan message.Message)

	z.Log.Info().Msg(`starting`)
	z.Loop.Run()
}

func (z *Zone) SubscribeOutgoing() chan message.Envelope {
	out := make(chan message.Envelope)
	go func(out chan message.Envelope) {
		for m := range z.Loop.Outgoing {
			out <- m
		}
	}(out)
	return out
}

func (z *Zone) SubscribeIncoming(in chan []byte) {
	go func(in chan []byte) {
		for m := range in {
			z.Log.Info().Str(`msg`, string(m)).Msg(`got message`)
			z.Loop.Incoming <-m
		}
	}(in)
}

func (z *Zone) Spawn(e core.Entity) {
	z.Loop.Spawn(e)
}