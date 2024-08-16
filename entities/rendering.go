package entities

import "github.com/hajimehoshi/ebiten/v2"

type Renderer interface {
	Draw(screen *ebiten.Image) error
}
