package stock_price

import (
	"fmt"
	"stocker-quant/util"
	"stocker-quant/web/app/model"

	"github.com/google/uuid"
)

// InsertStockPrice returns the rows that affected
// by insert statement
func InsertStockPrice(price StockPrice) (int64, uuid.UUID, error) {
	if !price.isValid() {
		return 0, uuid.UUID{}, fmt.Errorf("invalid parameter")
	}

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

	result, err := db.Exec(insertStatment)
	util.HandleError(err, "Insert Daily Stock Price Fail")
	if err != nil {
		return 0, uuid.UUID{}, err
	}

	rowsAffected, _ := result.RowsAffected()
	return rowsAffected, id, nil
}
