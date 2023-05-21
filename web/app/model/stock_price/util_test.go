package stock_price_test

import (
	"stocker-quant/configs"
	"stocker-quant/web/app/model"
	"stocker-quant/web/app/model/stock_price"

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

func insertTestRecord(price stock_price.StockPrice) uuid.UUID {
	const insertStatment = `
	INSERT INTO stock_price (
		id, company_id,
		update_date, price_date,
		open, close, high, low,
		change, change_percent,
		volume, amount
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
	);
	`

	db := model.GetDB()
	id := uuid.New()

	price.UpdateDate = price.UpdateDate.UTC()
	price.PriceDate = price.PriceDate.UTC()

	db.Exec(insertStatment,
		id,
		price.CompanyID, price.UpdateDate, price.PriceDate,
		price.Open, price.Close, price.High, price.Low,
		price.Change, price.ChangePercent,
		price.Volume, price.Amount,
	)

	return id
}
