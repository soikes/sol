package config

import (
	"io/ioutil"

	"soikke.li/sol/crdb"
	"soikke.li/sol/log"
	"soikke.li/sol/svc/game"
	"soikke.li/sol/svc/web"

	"gopkg.in/yaml.v2"
)

var defaultPath = `etc/local.yml`

type Config struct {
	Log log.Config `yaml:"log"`
	Sol struct {
		Web  web.Config  `yaml:"web"`
		Game game.Config `yaml:"game"`
	} `yaml:"sol"`
	Crdb crdb.Config `yaml:"crdb"`
}

func (cfg *Config) Load(path string) error {
	if path == `` {
		path = defaultPath
	}
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(dat, cfg)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *Config) Init() error {
	log := cfg.Log.Init()
	cfg.Crdb.Init(log)
	if err := cfg.Crdb.InitDB(); err != nil {
		log.Fatal().Err(err).Msg(`failed to initialize crdb`)
	}
	cfg.Sol.Game.Init(log)
	cfg.Sol.Web.Init(log)
	cfg.Sol.Web.InitDB(&cfg.Crdb)
	return nil
}

type Initializer interface {
	Init(log.Config)
}
