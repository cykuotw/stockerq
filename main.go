package main

import (
	"stocker-hf-data/configs"
	"stocker-hf-data/web/app"
)

func main() {
	configs.Init()

	app.Init()
	app.Setup()
	app.Run()
	defer app.End()
}
