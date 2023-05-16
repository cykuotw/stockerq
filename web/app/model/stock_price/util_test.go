package stock_price_test

import (
	"stocker-quant/configs"
	"stocker-quant/web/app/model"
)

func testPrep() {
	configs.Init()
	model.Connect()
}

func testClose() {
	model.Close()
}
