package components

import (
	"math"
	"time"

	"soikke.li/sol/primitives"
)

type Transform struct {
	Position primitives.Vec3 `json:"position"`
	Rotation primitives.Vec3 `json:"rotation"`
	Scale    primitives.Vec3 `json:"scale"`

	changed bool
}

func (t *Transform) Update(dt time.Duration) {
	t.Rotation.Clamp(0, 2*math.Pi)
	t.changed = false
	// log.Info().Float64(`X`, t.Position.X).Float64(`Y`, t.Position.Y).Float64(`Z`, t.Position.Z).Msg(`position update`)
	// log.Info().Float64(`X`, t.Rotation.X).Float64(`Y`, t.Rotation.Y).Float64(`Z`, t.Rotation.Z).Msg(`rotation update`)
}

func (t *Transform) AddPosition(v primitives.Vec3) {
	t.Position.Add(v)
	t.changed = true
}

func (t *Transform) Rotate(r primitives.Vec3) {
	t.Rotation.Add(r)
	t.changed = true
}
