package app

import "stockerq/web/app/model"

func Init() {
	model.Connect()
}

func Setup() {
	// do setup
}

func Run() {
	model.TestQuery()
}

func End() {
	model.Close()
}
