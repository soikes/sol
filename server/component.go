package sol

import (
	"soikke.li/sol/log"
)

type Component struct {
	Log log.Logger
}

func (c *Component) Init(name string, log log.Logger) {
	c.Log = log.WithComponent(name)
}