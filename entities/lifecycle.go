package entities

import "github.com/hajimehoshi/ebiten/v2"

type GameObjectLifecycle struct {
	HookDraw   func(self AbstractGameObjecter, screen *ebiten.Image) error
	HookUpdate func(self AbstractGameObjecter, deltaTime float64) error
}
