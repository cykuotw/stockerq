package stock_price

import (
	"fmt"
	"stocker-quant/util"
	"stocker-quant/web/app/model"

	"github.com/google/uuid"
)

const insertStatment = `
INSERT INTO stock_price (
	id, 
	company_id,
	update_date,
	price_date,
	open,
	close,
	high,
	low,
	change,
	change_percent,
	volume,
	amount
) VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8, $9, 
	$10, $11, $12
)
`

// InsertStockPrice returns the rows that affected
// by insert statement
func InsertStockPrice(price StockPrice) (int64, uuid.UUID, error) {
	if !price.isValid() {
		return 0, uuid.UUID{}, fmt.Errorf("invalid parameter")
	}

	db := model.GetDB()
	id := uuid.New()
	result, err := db.Exec(insertStatment,
		id,
		price.CompanyID, price.UpdateDate, price.PriceDate,
		price.Open, price.Close, price.High, price.Low,
		price.Change, price.ChangePercent,
		price.Volume, price.Amount,
	)
	util.HandleError(err, "Insert Daily Stock Price Fail")
	if err != nil {
		return 0, uuid.UUID{}, err
	}

	rowsAffected, _ := result.RowsAffected()
	return rowsAffected, id, nil
}
