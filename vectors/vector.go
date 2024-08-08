package vectors

type Vector2D struct {
	X float64
	Y float64
}

func (v *Vector2D) Add(other *Vector2D) {
	v.X += other.X
	v.Y += other.Y
}

func (v *Vector2D) Sub(other *Vector2D) {
	v.X -= other.X
	v.Y -= other.Y
}
