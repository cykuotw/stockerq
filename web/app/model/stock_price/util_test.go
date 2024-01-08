package stock_price_test

import (
	"fmt"
	"stockerq/configs"
	"stockerq/web/app/model"
	"stockerq/web/app/model/stock_price"

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
	db := model.GetAdminDB()
	db.Exec("DELETE FROM stock_price WHERE uuid='" + id.String() + "';")
}

func insertTestRecord(price stock_price.StockPrice) uuid.UUID {
	db := model.GetDB()
	id := uuid.New()

	price.UpdateDate = price.UpdateDate.UTC()
	price.PriceDate = price.PriceDate.UTC()

	insertStatment := fmt.Sprintf(`INSERT INTO stock_price (
		uuid, company_id, update_date, price_date, open, close, high, low,
		price_change, change_percent, volume, amount
	) VALUES (
		'%s', '%s', 
		'%s', '%s', 
		%d, %d, %d, %d,
		%d, %d, %d, %d);`,
		id.String(), price.CompanyID,
		price.UpdateDate.Format("2006-01-02 15:04:05"),
		price.PriceDate.Format("2006-01-02"),
		price.Open, price.Close, price.High, price.Low,
		price.PriceChange, price.ChangePercent,
		price.Volume, price.Amount,
	)

	db.Exec(insertStatment)

	return id
}
