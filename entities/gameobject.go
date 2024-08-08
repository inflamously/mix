package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"mix/vectors"
)

type GameObject struct {
	texture   *ebiten.Image
	draw      *ebiten.DrawImageOptions
	Lifecycle *GameObjectLifecycle
	Transform *Transform
	Data      State
}

func (g *GameObject) Update() error {
	deltaTime := 1.0 / 60

	if g.Lifecycle != nil && g.Lifecycle.HookUpdate != nil {
		return g.Lifecycle.HookUpdate(g, deltaTime)
	}

	return nil
}

func (g *GameObject) Draw(screen *ebiten.Image) error {
	screen.DrawImage(g.texture, g.draw)

	var err error = nil
	if g.Lifecycle != nil && g.Lifecycle.HookDraw != nil {
		err = g.Lifecycle.HookDraw(g, screen)
	}

	g.draw.GeoM.Reset()
	g.draw.GeoM.Translate(g.Transform.Position.X, g.Transform.Position.Y)
	g.draw.GeoM.Translate(-48, -48)
	g.draw.GeoM.Rotate(g.Transform.Rotation.Angle())
	g.draw.GeoM.Translate(48, 48)

	return err
}

func New(texture *ebiten.Image, data State, lifecycle *GameObjectLifecycle) *GameObject {
	gameobject := &GameObject{
		texture,
		&ebiten.DrawImageOptions{},
		lifecycle,
		nil,
		data,
	}

	transform := Transform{
		self: gameobject,
		Position: vectors.Vector2D{
			X: 0,
			Y: 0,
		},
	}

	gameobject.Transform = &transform

	return gameobject
}
