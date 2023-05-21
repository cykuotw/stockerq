package stock_price

import (
	"fmt"
	"stocker-quant/util"
	"stocker-quant/web/app/model"
	"time"
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

func GetStockPrice(startDate time.Time, endDate time.Time) ([]StockPrice, error) {
	// error check
	if startDate.IsZero() || endDate.IsZero() {
		return nil, fmt.Errorf("parameter error: startDate and endDate are neither zero")
	}
	if startDate.Compare(endDate) != -1 {
		return nil, fmt.Errorf("parameter error: startDate much be ealier than endDate")
	}

	// variable declare
	db := model.GetDB()
	var result []StockPrice

	startDate = startDate.UTC()
	endDate = endDate.UTC()

	// find price between startDate and endDate
	strQuery := fmt.Sprintf(`
			SELECT * FROM stock_price 
			WHERE price_date BETWEEN '%d-%d-%d' AND '%d-%d-%d';`,
		startDate.Year(), startDate.Month(), startDate.Day(),
		endDate.Year(), endDate.Month(), endDate.Day(),
	)
	rows, err := db.Query(strQuery)
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

		result = append(result, tmpPrice)
	}

	return result, nil
}
