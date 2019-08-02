package main

import (
	"engo-hello/src/systems"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"image/color"
)

//scene
type msScene struct{}

func (ms *msScene) Preload() {
	err := engo.Files.Load("textures/city.png") //preload to pre-loader
	if err != nil {
		panic(err)
	}
}

func (ms *msScene) Setup(u engo.Updater) {
	world, _ := u.(*ecs.World)

	//set bg
	common.SetBackground(color.White)
	//add system
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&systems.CityBuildingSystem{})
	//register input
	engo.Input.RegisterButton("AddCity", engo.KeyF1)

	//make entity
	city := City{BasicEntity: ecs.NewBasic()}
	city.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 10, Y: 10},
		Width:    303,
		Height:   641,
	}
	texture, err := common.LoadedSprite("textures/city.png") //load from pre-loader
	if err != nil {
		panic(err)
	}
	city.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{X: 1, Y: 1},
	}

	//add entity to system
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&city.BasicEntity, &city.RenderComponent, &city.SpaceComponent)
		}
	}
}

func (ms *msScene) Type() string {
	return "myGame"
}

//entity
type City struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

func main() {
	opts := engo.RunOptions{
		Title:  "Hello Engo!",
		Width:  400,
		Height: 400,
	}
	scene := &msScene{}
	engo.Run(opts, scene)
}
