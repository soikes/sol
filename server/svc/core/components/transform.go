package components

import (
	"math"
	"time"

	"soikke.li/sol/primitives"
)

type Transform struct {
	Position primitives.Vec3
	Rotation primitives.Vec3
	Scale    primitives.Vec3

	changed bool
}

func (t *Transform) Update(dt time.Duration) {
	t.Rotation.Clamp(0, 2*math.Pi)
	t.changed = false
}

func (t *Transform) AddPosition(v primitives.Vec3) {
	t.Position.Add(v)
	t.changed = true
}

func (t *Transform) Rotate(r primitives.Vec3) {
	t.Rotation.Add(r)
	t.changed = true
}