package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type CityBuildingSystem struct {
	world *ecs.World
}

func (cbs *CityBuildingSystem) Remove(ecs.BasicEntity) {}

func (cbs *CityBuildingSystem) Update(dt float32) {
	if engo.Input.Button("AddCity").JustPressed() {
		//entity
		type City struct {
			ecs.BasicEntity
			common.SpaceComponent
			common.RenderComponent
		}
		//make entity
		city := City{}
		// BasicEntity(id)
		city.BasicEntity = ecs.NewBasic()
		// SpaceComponent(like collider)
		city.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{X: engo.Input.Mouse.X, Y: engo.Input.Mouse.Y},
			Width:    30,
			Height:   64,
		}
		// RenderComponent(render)
		texture, err := common.LoadedSprite("textures/city.png") //load from pre-loader
		if err != nil {
			panic(err)
		}
		city.RenderComponent = common.RenderComponent{
			Drawable: texture,
			Scale:    engo.Point{X: 0.1, Y: 0.1},
		}

		//add entity to system
		for _, system := range cbs.world.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				sys.Add(&city.BasicEntity, &city.RenderComponent, &city.SpaceComponent)
			}
		}
	}
}

func (cbs *CityBuildingSystem) New(w *ecs.World) {
	cbs.world = w
	//register input
	engo.Input.RegisterButton("AddCity", engo.KeyF1)
}
