package main

import (
	"engo-hello/src/systems"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"image"
	"image/color"
)

//scene
type msScene struct{}

func (ms *msScene) Preload() {
	//preload file to pre-loader
	err := engo.Files.Load("textures/city.png", "tilemap/TrafficMap.tmx")
	if err != nil {
		panic(err)
	}
}

func (ms *msScene) Setup(u engo.Updater) {
	//get world
	world, _ := u.(*ecs.World)

	//set bg
	common.SetBackground(color.White)
	//add system
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(common.NewKeyboardScroller(40, engo.DefaultHorizontalAxis, engo.DefaultVerticalAxis))
	//world.AddSystem(&common.EdgeScroller{100, 20})
	world.AddSystem(&common.MouseZoomer{-0.125})
	world.AddSystem(&systems.CityBuildingSystem{})

	//HUD entity
	hud := HUD{BasicEntity: ecs.NewBasic()}
	hud.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{0, engo.WindowHeight() - 30},
		Width:    30,
		Height:   30,
	}
	imageUniform := image.NewUniform(color.NRGBA{205, 205, 205, 0x6f})
	imageNRGBA := common.ImageToNRGBA(imageUniform, 200, 200)
	hudImage := common.NewImageObject(imageNRGBA)
	hudTexture := common.NewTextureSingle(hudImage)
	hud.RenderComponent = common.RenderComponent{
		Drawable: hudTexture,
		Scale:    engo.Point{},
		Repeat:   common.Repeat,
	}
	hud.RenderComponent.SetShader(common.HUDShader)
	hud.RenderComponent.SetZIndex(1)
	//add entity to system
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&hud.BasicEntity, &hud.RenderComponent, &hud.SpaceComponent)
		}
	}
}

func (ms *msScene) Type() string {
	return "myGame"
}

type HUD struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

func main() {
	opts := engo.RunOptions{
		Title:          "Hello Engo!",
		Width:          100,
		Height:         100,
		StandardInputs: true,
		NotResizable:   true,
	}
	scene := &msScene{}
	engo.Run(opts, scene)
}
