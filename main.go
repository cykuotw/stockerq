package main

import (
	"stocker-quant/configs"
	"stocker-quant/web/app"
)

func main() {
	configs.Init()

	app.Init()
	app.Setup()
	app.Run()
	defer app.End()
}
