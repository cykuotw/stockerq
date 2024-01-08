package main

import (
	"stockerq/configs"
	"stockerq/web/app"
)

func main() {
	configs.Init()

	app.Init()
	app.Setup()
	app.Run()
	defer app.End()
}
