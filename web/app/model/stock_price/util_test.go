package stock_price_test

import (
	"stocker-quant/configs"
	"stocker-quant/web/app/model"

	"github.com/google/uuid"
)

func testPrep() {
	configs.Init()
	model.Connect()
}

func testClose() {
	model.Close()
}

func deleteTestRecord(id uuid.UUID) {
	db := model.GetDB()
	db.Exec("DELETE FROM stock_price WHERE id='" + id.String() + "'")
}
