package sol

import (
	"io/ioutil"

	"soikke.li/sol/svc/web"
	"soikke.li/sol/crdb"

	"gopkg.in/yaml.v2"
)

var defaultPath = `etc/local.yml`

type Config struct {
	Sol struct {
		Web web.Config `yaml:web`
	} `yaml:sol`
	Crdb crdb.Config `yaml:crdb`
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