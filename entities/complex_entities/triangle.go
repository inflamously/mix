package complex_entities

/**
 * This class implements and represents a complex game object that has its own logic and rendering stuff.
 */

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"mix/entities"
	"mix/render"
)

type TriangleState struct {
	image  *ebiten.Image
	points [][]float64
}

type Triangle struct {
	entities.AbstractGameObject
	Lifecycle entities.GameObjectLifecycle
	State     entities.State
}

func (t *Triangle) Initialize() error {
	t.State.Data = TriangleState{
		points: [][]float64{
			{0.0, 0.0},
			{0.0, 0.0},
			{0.0, 0.0},
		},
	}

	state := t.GetState()
	state.image = ebiten.NewImage(128, 128)

	render.DrawRectangle(state.image, image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{1, 1},
	}, render.ColorPalette[render.COLOR_BLUE])

	return nil
}

func (t *Triangle) GetState() TriangleState {
	return t.State.Data.(TriangleState)
}

func (t *Triangle) Update(game ebiten.Game) error {
	//deltaTime := 1.0 / 60

	return nil
}

func (t *Triangle) Draw(screen *ebiten.Image) error {
	return nil
}
