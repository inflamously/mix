package rotations

type Radians = float64
type Degree = float64

const Angle2Radians = 0.01745329251994329576923690768489

type Rotation struct {
	angle Radians
}

func (r *Rotation) Angle() Degree {
	return r.angle / Angle2Radians
}

func (r *Rotation) SetRotation(rotation Degree) {
	r.angle = rotation
}
