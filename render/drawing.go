package render

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
)

type ColorFromPalette = int

const (
	COLOR_UNKNOWN = iota
	COLOR_BLUE
)

var ColorPalette = map[ColorFromPalette]color.RGBA{
	COLOR_BLUE: color.RGBA{
		R: 0,
		G: 0,
		B: 255,
		A: 255,
	},
}

func DrawRectangle(image *ebiten.Image, rect image.Rectangle, color color.RGBA) {
	vector.DrawFilledRect(image, float32(rect.Min.X), float32(rect.Min.Y), float32(rect.Max.X), float32(rect.Max.Y), color, true)
}
