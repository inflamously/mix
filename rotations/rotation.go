package rotations

type Radians = float64
type Degree = float64

const HalfCircleRadians = 0.01745329251994329576923690768489

type Rotation struct {
	angle Radians
}

func (r *Rotation) Angle() Degree {
	return r.angle / HalfCircleRadians
}

func (r *Rotation) SetRotation(rotation Degree) {
	r.angle = rotation
}
