package primitives

type Vector interface {
	Multiply(float64)
	Add(Vector)
	AddScalar(float64)
	Dot(Vector) float64
	Cross(Vector) Vector
	Magnitude() float64
}

type Vec2D struct {
	X float64
	Y float64
}

func (v *Vec2D) Multiply(in float64) {
	v.X *= in
	v.Y *= in
}

func (v *Vec2D) Add(in Vec2D) {
	v.X += in.X
	v.Y += in.Y
}

func (v *Vec2D) AddScalar(in float64) {
	v.X += in
	v.Y += in
}

func (v *Vec2D) Dot(in Vec2D) float64 {
	return (v.X * in.X) + (v.Y * in.Y)
}

func (v *Vec2D) Cross(in Vec2D) Vec2D { return Vec2D{} } //N/A

type Vec3D struct {
	X float64
	Y float64
	Z float64
}

func (v *Vec3D) Multiply(in float64) {
	v.X *= in
	v.Y *= in
	v.Z *= in
}
