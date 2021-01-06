package components

import (
	"math"
	"time"

	// "time"

	"github.com/rs/zerolog/log"
	"soikke.li/sol/primitives"
	// "soikke.li/sol/svc/core"
)

type Physics struct {
	Velocity primitives.Vec3 `json:"velocity"`
	MaxSpeed float64         `json:"-"`

	AccelerationFactor float64 `json:"-"`
	accelerating       bool    `json:"-"`

	RotationFactor    float64         `json:"-"`
	rotating          bool            `json:"-"`
	RotationDirection primitives.Vec3 `json:"rotationDirection"`
	RotationBuf       primitives.Vec3 `json:"-"`

	Transform *Transform `json:"-"`
}

func NewPhysics(t *Transform) Physics {
	return Physics{Transform: t}
}

func (p *Physics) Update(dt time.Duration) {
	p.ApplyRotation()
	p.CalculateVelocity(dt)
	p.UpdatePosition(dt)
}

func (p *Physics) ApplyRotation() {
	if p.rotating {
		p.RotationBuf = p.RotationDirection
		p.RotationBuf.MultScalar(p.RotationFactor)
		p.Transform.Rotate(p.RotationBuf)
	}
}

func (p *Physics) CalculateVelocity(dt time.Duration) {
	if p.accelerating {
		vix := p.Velocity.X
		vfx := vix + (p.AccelerationFactor*math.Sin(p.Transform.Rotation.Y))*dt.Seconds()

		viz := p.Velocity.Z
		vfz := viz + (p.AccelerationFactor*math.Cos(p.Transform.Rotation.Y))*dt.Seconds()

		dir := primitives.Vec2{X: vfx, Y: vfz}
		mag := dir.Magnitude()
		dir.Normalize()

		if mag <= p.MaxSpeed {
			p.Velocity.X = vfx
			p.Velocity.Z = vfz
		} else {
			p.Velocity.X = dir.X * p.MaxSpeed
			p.Velocity.Z = dir.Y * p.MaxSpeed
		}
	}
}

func (p *Physics) UpdatePosition(dt time.Duration) {
	d := primitives.Vec3{X: p.Velocity.X, Z: p.Velocity.Z}
	d.MultScalar(dt.Seconds())
	p.Transform.AddPosition(d)
}

func (p *Physics) Accelerate() {
	log.Info().Msg(`accelerating!`)
	p.accelerating = true
}

func (p *Physics) StopAccelerating() {
	p.accelerating = false
}

func (p *Physics) RotateY(r float64) {
	p.rotating = true
	p.RotationDirection.Y = r
}

func (p *Physics) StopRotating() {
	p.rotating = false
	p.RotationDirection = primitives.Vec3{}
}
