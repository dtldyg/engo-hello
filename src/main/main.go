package main

import (
	"github.com/EngoEngine/engo"
	_ "github.com/EngoEngine/engo/common"
)

type msScene struct{}

func (ms *msScene) Preload() {
	engo.Files.Load("textures/city.png")
}

func (ms *msScene) Setup(engo.Updater) {}

func (ms *msScene) Type() string {
	return "myGame"
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
