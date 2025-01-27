package core

import (
	"time"
)

type Component interface {
	Update(time.Duration)
	Attach(string)
}
