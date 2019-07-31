package main

import (
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
	common.SetBackground(color.White)
	world, _ := u.(*ecs.World)

	//add system to world
	world.AddSystem(&common.RenderSystem{})

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
	ecs.BasicEntity        //继承基础实体
	common.SpaceComponent  //空间组件：在哪渲染
	common.RenderComponent //渲染组件：渲染什么
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
