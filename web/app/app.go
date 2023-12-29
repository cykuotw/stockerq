package app

import "stocker-hf-data/web/app/model"

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
