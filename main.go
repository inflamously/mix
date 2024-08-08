package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"mix/entities"
	"mix/resources"
)

type Game struct {
	initialized bool
	Gameobjects []*entities.GameObject
}

func (g *Game) Initialize() {
	g.initialized = true

	img, err := resources.LoadImage("./assets/cube.png")
	if err != nil {
		return
	}

	cube := entities.New(img, entities.State{
		Data: &struct {
			name string
		}{
			name: "Hello",
		},
	}, &entities.GameObjectLifecycle{
		HookUpdate: func(self *entities.GameObject, deltaTime float64) error {
			self.Transform.SetPosition(32, 32)
			self.Transform.SetRotation(self.Transform.Rotation.Angle() + 10*deltaTime)
			fmt.Printf("%v:%v\n", self.Transform.Bounds().Dx(), self.Transform.Bounds().Dy())
			return nil
		},
		HookDraw: func(self *entities.GameObject, screen *ebiten.Image) error {
			return nil
		},
	})

	g.Gameobjects = append(g.Gameobjects, cube)
}

func (g *Game) Update() error {
	if !g.initialized {
		g.Initialize()
	}

	for i := range g.Gameobjects {
		gameobject := g.Gameobjects[i]
		err := gameobject.Update()
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := range g.Gameobjects {
		gameobject := g.Gameobjects[i]
		err := gameobject.Draw(screen)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Mix")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
