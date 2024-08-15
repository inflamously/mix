package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
	"mix/entities"
	"mix/entities/complex_entities"
)

type Screen struct {
	Size image.Point
}

type Game struct {
	Initialized bool
	Screen      Screen
	GameObjects map[string]entities.AbstractGameObject
}

func (g *Game) registerGameObjects() map[string]entities.AbstractGameObject {
	//img, err := resources.LoadImage("./assets/cube.png")
	//if err != nil {
	//	return
	//}
	//

	//cube := entities.New(img, entities.State{
	//	Data: &struct {
	//		name string
	//	}{
	//		name: "Hello",
	//	},
	//}, &entities.GameObjectLifecycle{
	//	HookUpdate: func(self *entities.GameObject, deltaTime float64) error {
	//		self.Transform.SetPosition(32, 32)
	//		self.Transform.SetRotation(self.Transform.Rotation.Angle() + deltaTime)
	//		fmt.Printf("%v:%v\n", self.Transform.Bounds().Dx(), self.Transform.Bounds().Dy())
	//		return nil
	//	},
	//	HookDraw: func(self *entities.GameObject, screen *ebiten.Image) error {
	//		return nil
	//	},
	//})

	//g.GameObjects = append(g.GameObjects, cube)

	//img := ebiten.NewImage(128, 128)
	//blue := color.RGBA{
	//	R: 255,
	//	G: 0,
	//	B: 255,
	//	A: 255,
	//}

	//points := [][]float32{
	//	{64 / 2, 0},
	//	{0, 64},
	//	{64, 64},
	//}
	//
	//vector.DrawFilledRect(img, points[0][0], points[0][1], 2, 2, blue, true)
	//vector.DrawFilledRect(img, points[1][0], points[1][1], 2, 2, blue, true)
	//vector.DrawFilledRect(img, points[2][0], points[2][1], 2, 2, blue, true)
	//triangle := entities.New(img, entities.State{}, &entities.GameObjectLifecycle{
	//	HookUpdate: nil,
	//	HookDraw:   nil,
	//})
	//triangle.Transform.SetPosition(float64(g.Screen.Size.X/2)-64, float64(g.Screen.Size.Y/2)-64)

	nTriangle := complex_entities.Triangle{
		State: entities.State{},
	}

	return map[string]entities.AbstractGameObject{
		//"triangle":  triangle,
		"nTriangle": &nTriangle,
	}
}

func (g *Game) Initialize() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("initialization failed.")
			log.Fatal(r)
		}
	}()

	g.GameObjects = g.registerGameObjects()
	g.Initialized = true
	for k := range g.GameObjects {
		fmt.Printf("initialize gameobject '%v'\n", k)
		err := g.GameObjects[k].Initialize()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Game) Update() error {
	if !g.Initialized {
		g.Initialize()
	}

	for i := range g.GameObjects {
		gameObject := g.GameObjects[i]
		err := gameObject.Update(g)
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := range g.GameObjects {
		gameObject := g.GameObjects[i]
		err := gameObject.Draw(screen)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Screen.Size.X, g.Screen.Size.Y
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Mix")
	if err := ebiten.RunGame(&Game{
		Screen: Screen{Size: image.Point{X: 640, Y: 480}},
	}); err != nil {
		log.Fatal(err)
	}
}
