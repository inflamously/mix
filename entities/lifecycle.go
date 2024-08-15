package entities

import "github.com/hajimehoshi/ebiten/v2"

type GameObjectLifecycle struct {
	HookDraw   func(self AbstractGameObject, screen *ebiten.Image) error
	HookUpdate func(self AbstractGameObject, deltaTime float64) error
}
