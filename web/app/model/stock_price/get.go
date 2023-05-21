package stock_price

import (
	"stocker-quant/util"
	"stocker-quant/web/app/model"
)

func GetStockPriceLatest() ([]StockPrice, error) {
	// variables declare
	db := model.GetDB()
	var results []StockPrice

	// find records with latest date
	rows, err := db.Query(`
				SELECT * FROM stock_price WHERE price_date = 
				(
					SELECT price_date 
					FROM stock_price 
					GROUP BY price_date
					ORDER BY price_date DESC
					LIMIT 1
				);`)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmpPrice StockPrice
		err := rows.Scan(&tmpPrice.Id, &tmpPrice.CompanyID,
			&tmpPrice.UpdateDate, &tmpPrice.PriceDate,
			&tmpPrice.Open, &tmpPrice.Close, &tmpPrice.High, &tmpPrice.Low,
			&tmpPrice.Change, &tmpPrice.ChangePercent,
			&tmpPrice.Volume, &tmpPrice.Amount)
		util.HandleError(err, "Scan rows to StockPrice struct Fail")
		if err != nil {
			continue
		}

		results = append(results, tmpPrice)
	}

	return results, nil
}
