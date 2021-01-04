package game

import (
	"soikke.li/sol"
	"soikke.li/sol/log"
	"soikke.li/sol/svc/core"
	"soikke.li/sol/svc/core/loop"
)

type Zone struct {
	sol.Component

	Id   string
	Loop loop.Config
	Hub  Hub
}

func (z *Zone) Init(log log.Logger) {
	z.Component.Init(`zone`, log)

	z.Loop = loop.Config{}
	z.Loop.Init(log)

	z.Log.Info().Msg(`starting`)
	z.Loop.Run()

	z.Hub = NewHub()

	z.SubscribeIncoming(z.Hub.Incoming)
	z.SubscribeOutgoing(z.Hub.Broadcast)
}

func (z *Zone) SubscribeOutgoing(out chan []byte) {
	go func(out chan []byte) {
		for m := range z.Loop.Outgoing {
			out <- m
		}
	}(out)
}

func (z *Zone) SubscribeIncoming(in chan []byte) {
	go func(in chan []byte) {
		for m := range in {
			z.Log.Info().Str(`msg`, string(m)).Msg(`got message`)
			z.Loop.Incoming <- m
		}
	}(in)
}

func (z *Zone) Spawn(e core.Entity) {
	z.Loop.Spawn(e)
}

func (z *Zone) RegisterClient(c *Client) {
	z.Hub.RegisterClient(c)
}

func (z *Zone) DeregisterClient(c *Client) {
	z.Hub.DeregisterClient(c)
}

var theZone *Zone //TODO use a singleton to debug until we can construct the world with many zones

func GetZone() *Zone {
	if theZone == nil {
		theZone = &Zone{}
		theZone.Init(log.Logger{})
	}
	return theZone
}
