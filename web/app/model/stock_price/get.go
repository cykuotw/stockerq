package stock_price

import (
	"fmt"
	apperror "stockerq/web/app/app-error"
	"stockerq/web/app/model"
	"time"
)

func GetStockPriceLatest() ([]StockPrice, *apperror.ModelError) {
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
		return nil, apperror.NewModelError(err)
	}

	for rows.Next() {
		var id uint64
		var priceDate string
		var updateDate string
		var tmpPrice StockPrice
		err := rows.Scan(&id,
			&tmpPrice.Uuid, &tmpPrice.CompanyID,
			&updateDate, &priceDate,
			&tmpPrice.Open, &tmpPrice.Close, &tmpPrice.High, &tmpPrice.Low,
			&tmpPrice.PriceChange, &tmpPrice.ChangePercent,
			&tmpPrice.Volume, &tmpPrice.Amount)
		tmpPrice.PriceDate, err = time.Parse("2006-01-02", priceDate)
		tmpPrice.UpdateDate, err = time.Parse("2006-01-02 15:04:05", updateDate)

		if err != nil {
			return nil, apperror.NewModelError(err)
		}

		results = append(results, tmpPrice)
	}

	return results, nil
}

func GetStockPrice(startDate time.Time, endDate time.Time) ([]StockPrice, *apperror.ModelError) {
	if startDate.IsZero() || endDate.IsZero() {
		return nil, apperror.NewModelError(apperror.ErrZeroDate)
	}
	diff := endDate.Sub(startDate)
	if diff < 0 {
		return nil, apperror.NewModelError(apperror.ErrReverseDate)
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
		return nil, apperror.NewModelError(err)
	}

	for rows.Next() {
		var id uint64
		var priceDate string
		var updateDate string
		var tmpPrice StockPrice
		err := rows.Scan(&id,
			&tmpPrice.Uuid, &tmpPrice.CompanyID,
			&updateDate, &priceDate,
			&tmpPrice.Open, &tmpPrice.Close, &tmpPrice.High, &tmpPrice.Low,
			&tmpPrice.PriceChange, &tmpPrice.ChangePercent,
			&tmpPrice.Volume, &tmpPrice.Amount)
		tmpPrice.PriceDate, err = time.Parse("2006-01-02", priceDate)
		tmpPrice.UpdateDate, err = time.Parse("2006-01-02 15:04:05", updateDate)

		if err != nil {
			return nil, apperror.NewModelError(err)
		}

		result = append(result, tmpPrice)
	}

	return result, nil
}
