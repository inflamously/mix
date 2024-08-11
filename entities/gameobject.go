package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"mix/vectors"
)

type Updater interface {
	Update() error
}

type Initializer interface {
	Initialize() error
}

type Renderer interface {
	Draw(screen *ebiten.Image) error
}

type AbstractGameObjecter interface {
	Initializer
	Updater
	Renderer
}

type GameObject struct {
	texture   *ebiten.Image
	draw      *ebiten.DrawImageOptions
	Lifecycle *GameObjectLifecycle
	Transform *Transform
	Data      State
}

func (g *GameObject) Initialize() error {
	return nil
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

	// Reset to initials
	g.draw.GeoM.Reset()

	// Prepare rotation
	g.rotateAroundCenter()

	// Translate to position
	g.draw.GeoM.Translate(g.Transform.Position.X, g.Transform.Position.Y)

	return err
}

func (g *GameObject) rotateAroundCenter() {
	bounds := g.Transform.Bounds().Size()
	g.draw.GeoM.Translate(float64(-bounds.X/2), float64(-bounds.Y/2))
	g.draw.GeoM.Rotate(g.Transform.Rotation.Angle())
	g.draw.GeoM.Translate(float64(bounds.X/2), float64(bounds.Y/2))
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
