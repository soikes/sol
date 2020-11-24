package loop

import (
	"sync"
	"time"

	"soikke.li/sol"
	"soikke.li/sol/log"
	"soikke.li/sol/message"
	"soikke.li/sol/svc/core"
)

type Config struct {
	sol.Component

	Rate time.Duration `yaml:"rate"`

	Entities map[string]core.Entity
	mu *sync.Mutex

	Incoming chan []byte // TODO Consider how to structure this per client instead of mux over one channel
	Outgoing chan message.Envelope
}

func (cfg *Config) Init(log log.Logger) {
	cfg.Component.Init(`loop`, log)

	cfg.Rate = time.Second

	if cfg.Rate == 0 {
		cfg.Log.Fatal().Msg(`missing rate, cannot start game loop`)
	}
	cfg.mu = &sync.Mutex{}

	cfg.Entities = map[string]core.Entity{}
	cfg.Incoming = make(chan []byte)
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
			// cfg.Log.Info().Str(`msg`, string(msg)).Msg(`got message`)
			i, err := message.Unmarshal(msg)
			if err != nil {
				cfg.Log.Info().Err(err).Msg(`unknown message type`)
			}
			e := cfg.Entities[i.ID]
			cfg.Log.Info().Str(`id`, e.Id).Msg(`input message for component`)
		}
	}()
} 

func (cfg *Config) Spawn(e core.Entity) {
	cfg.Log.Info().Str(`id`, e.Id).Msg(`spawning entity`)
	cfg.mu.Lock()
	cfg.Entities[e.Id] = e
	cfg.mu.Unlock()
}