package entities

import "github.com/hajimehoshi/ebiten/v2"

type GameInstance interface {
	Initialize()
	Update() error
	Draw(screen *ebiten.Image)
}

type Initializer interface {
	Initialize(game GameInstance) error
}

type Updater interface {
	Update(game GameInstance) error
}

type Renderer interface {
	Draw(screen *ebiten.Image) error
}

type GameObjectLifecycle struct {
	HookDraw   func(self AbstractGameObject, screen *ebiten.Image) error
	HookUpdate func(self AbstractGameObject, deltaTime float64) error
}
