package primitives

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func (v *Vec2) Normalize() {
	m := v.Magnitude()
	if m > 0 {
		v.X = v.X / m
		v.Y = v.Y / m
	}
}

func (v *Vec2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}