package entities

import "github.com/hajimehoshi/ebiten/v2"

type GameObjectLifecycle struct {
	HookDraw   func(self *GameObject, screen *ebiten.Image) error
	HookUpdate func(self *GameObject, deltaTime float64) error
}
