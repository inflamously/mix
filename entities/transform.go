package entities

import (
	"image"
	"math"
	"mix/rotations"
	"mix/vectors"
)

type Transform struct {
	self     *GameObject
	Position vectors.Vector2D
	Rotation rotations.Rotation
}

func (g *Transform) SetPosition(x, y float64) {
	g.Position = vectors.Vector2D{X: x, Y: y}
}

func (g *Transform) SetRotation(degree float64) {
	g.Rotation.SetRotation((math.Mod(degree, 359)) * rotations.HalfCircleRadians)
}

func (g *Transform) Bounds() image.Rectangle {
	return g.self.texture.Bounds()
}
