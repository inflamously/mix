package complex_entities

/**
 * This class implements and represents a complex game object that has its own logic and rendering stuff.
 */

import (
	"github.com/hajimehoshi/ebiten/v2"
	"mix/entities"
)

type TriangleState struct {
	points [][]float64
}

type Triangle struct {
	entities.AbstractGameObjecter
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

	return nil
}

func (t *Triangle) Update() error {
	return t.Lifecycle.HookUpdate(t, 1.0/60)
}

func (t *Triangle) Draw(screen *ebiten.Image) error {
	return t.Lifecycle.HookDraw(t, screen)
}
