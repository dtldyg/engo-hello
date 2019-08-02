package systems

import (
	"fmt"
	"github.com/EngoEngine/ecs"
)

type CityBuildingSystem struct{}

func (*CityBuildingSystem) Remove(ecs.BasicEntity) {}

func (*CityBuildingSystem) Update(dt float32) {}

func (*CityBuildingSystem) New(*ecs.World) {
	fmt.Println("CityBuildingSystem was added to the Scene")
}
