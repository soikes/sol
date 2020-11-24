package primitives

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v *Vec3) MultScalar(s float64) {
	v.X = v.X * s
	v.Y = v.Y * s
	v.Z = v.Z * s
}

func (v *Vec3) Add(a Vec3) {
	v.X += a.X
	v.Y += a.Y
	v.Z += a.Z
}

func (v *Vec3) Clamp(min, max float64) {
	if v.X < min {
		v.X = max
	}
	if v.X > max {
		v.X = min
	}

	if v.Y < min {
		v.Y = max
	}
	if v.Y > max {
		v.Y = min
	}

	if v.Z < min {
		v.Z = max
	}
	if v.Z > max {
		v.Z = min
	}
}