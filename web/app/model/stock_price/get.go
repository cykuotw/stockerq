package stock_price

import (
	"fmt"
	"stocker-quant/util"
	"stocker-quant/web/app/model"
	"time"
)

func GetStockPriceLatest() ([]StockPrice, time.Time, error) {
	// variables declare
	var latestDateUTC time.Time
	db := model.GetDB()
	var results []StockPrice

	// find latest date
	err := db.QueryRow(`
				SELECT price_date 
				FROM stock_price 
				GROUP BY price_date
				ORDER BY price_date DESC
				LIMIT 1;`).Scan(&latestDateUTC)

	util.HandleError(err, "Get Latest Price Date Fail")
	if err != nil {
		return nil, time.Time{}, err
	}

	// find records with latest date
	strQuery := fmt.Sprintf(`SELECT * FROM stock_price WHERE price_date = '%d-%d-%d';`,
		latestDateUTC.Year(), latestDateUTC.Month(), latestDateUTC.Day())
	rows, err := db.Query(strQuery)
	defer rows.Close()
	if err != nil {
		return nil, time.Time{}, err
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

	return results, latestDateUTC, nil
}
