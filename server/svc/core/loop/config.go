package loop

import (
	"sync"
	"time"

	"soikke.li/sol"
	"soikke.li/sol/log"
	"soikke.li/sol/message"
	"soikke.li/sol/svc/core"
	"soikke.li/sol/svc/core/components"
)

type Config struct {
	sol.Component

	Rate time.Duration `yaml:"rate"`

	Entities map[string]core.Entity // TODO use slice for performance here
	mu       *sync.Mutex

	Incoming chan []byte // TODO Consider how to structure this per client instead of mux over one channel
	Outgoing chan []byte
}

func (cfg *Config) Init(log log.Logger) {
	cfg.Component.Init(`loop`, log)

	cfg.Rate = time.Second / 30

	if cfg.Rate == 0 {
		cfg.Log.Fatal().Msg(`missing rate, cannot start game loop`)
	}
	cfg.mu = &sync.Mutex{}

	cfg.Entities = map[string]core.Entity{}
	cfg.Incoming = make(chan []byte)
	cfg.Outgoing = make(chan []byte)
}

func (cfg *Config) Run() error {
	go func() {
		t := time.NewTicker(cfg.Rate)
		defer t.Stop()
		last := time.Now()

		for cur := range t.C {
			dt := cur.Sub(last)
			cfg.Log.Info().Dur(`dt`, dt).Msg(`elapsed`)
			for _, e := range cfg.Entities {
				cfg.Log.Info().Str(`id`, e.Id).Msg(`updating entity`)
				e.Update(dt)
				for _, c := range e.Components {
					t, ok := c.(*components.Transform)
					if ok {
						msg := message.Transform{
							ID: e.Id,
							X:  t.Position.X,
							Y:  t.Position.Y,
							Z:  t.Position.Z,
						}
						b, err := msg.Marshal()
						if err != nil {
							cfg.Log.Error().Err(err).Msg(`failed to marshal outgoing transform message`)
						}
						cfg.Outgoing <- b
						// select {
						// case cfg.Outgoing <- b:
						// 	cfg.Log.Info().Msg(`sent transform message`)
						// default:
						// 	cfg.Log.Info().Msg(`skipped outgoing transform message`)
						// }
					}
				}
			}
			last = cur
		}
	}()

	cfg.CollectInputs()

	return nil
}

func (cfg *Config) CollectInputs() {
	// cfg.Log.Info().Msg(`collecting inputs`)
	go func() {
		for msg := range cfg.Incoming {
			cfg.Log.Info().Str(`msg`, string(msg)).Msg(`got message in loop`)
			var i message.Input
			err := message.Unmarshal(msg, &i)
			if err != nil {
				cfg.Log.Info().Err(err).Msg(`unknown message type`)
			}
			e := cfg.Entities[i.ID]
			for _, c := range e.Components {
				in, ok := c.(*components.Input)
				if ok {
					cfg.Log.Info().Str(`id`, e.Id).Msg(`input message for component`)
					in.QueueInput(i)
				}
			}
		}
	}()
}

func (cfg *Config) Spawn(e core.Entity) {
	cfg.Log.Info().Str(`id`, e.Id).Msg(`spawning entity`)
	cfg.mu.Lock()
	cfg.Entities[e.Id] = e
	cfg.mu.Unlock()
}
