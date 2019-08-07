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

	//set default bg
	common.SetBackground(color.White)
	//add system
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(common.NewKeyboardScroller(200, engo.DefaultHorizontalAxis, engo.DefaultVerticalAxis))
	//world.AddSystem(&common.EdgeScroller{100, 20})
	world.AddSystem(&common.MouseZoomer{-0.125})
	world.AddSystem(&systems.CityBuildingSystem{})

	//set map
	if err := setMap(world, "tilemap/TrafficMap.tmx"); err != nil {
		panic(err)
	}

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

func setMap(world *ecs.World, url string) error {
	//set map
	resource, err := engo.Files.Resource(url)
	if err != nil {
		return err
	}
	tmxResource := resource.(common.TMXResource)
	levelData := tmxResource.Level
	tiles := make([]*Tile, 0)
	for _, tileLayer := range levelData.TileLayers {
		for _, tileElement := range tileLayer.Tiles {
			if tileElement.Image != nil {
				tile := &Tile{BasicEntity: ecs.NewBasic()}
				tile.RenderComponent = common.RenderComponent{
					Drawable: tileElement,
					Scale:    engo.Point{1, 1},
				}
				tile.SpaceComponent = common.SpaceComponent{
					Position: tileElement.Point,
					Width:    0,
					Height:   0,
				}
				tiles = append(tiles, tile)
			}
		}
	}
	// add the tiles to the RenderSystem
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, v := range tiles {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		}
	}
	common.CameraBounds = levelData.Bounds()
	return nil
}

func (ms *msScene) Type() string {
	return "myGame"
}

type HUD struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

type Tile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
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
