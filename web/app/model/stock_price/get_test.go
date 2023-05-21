package stock_price_test

import (
	stock "stocker-quant/web/app/model/stock_price"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetStockPrice(t *testing.T) {
	testPrep()
	defer testClose()

	loc, _ := time.LoadLocation("Asia/Taipei")
	validPrice := []stock.StockPrice{
		{
			CompanyID:     "2330",
			UpdateDate:    time.Now().In(loc),
			PriceDate:     time.Now().In(loc),
			Open:          50800,
			Close:         50300,
			High:          50800,
			Low:           50000,
			Change:        -700,
			ChangePercent: -139,
			Volume:        19385820,
			Amount:        9753130414,
		},
		{
			CompanyID:     "2454",
			UpdateDate:    time.Now().In(loc),
			PriceDate:     time.Now().In(loc),
			Open:          69100,
			Close:         69800,
			High:          69800,
			Low:           69000,
			Change:        900,
			ChangePercent: 131,
			Volume:        5282896,
			Amount:        3673257462,
		},
		{
			CompanyID:     "2603",
			UpdateDate:    time.Now().In(loc),
			PriceDate:     time.Now().In(loc).AddDate(0, 0, -1),
			Open:          15100,
			Close:         15050,
			High:          15200,
			Low:           15000,
			Change:        -50,
			ChangePercent: -33,
			Volume:        15427176,
			Amount:        2327960531,
		},
		{
			CompanyID:     "2891",
			UpdateDate:    time.Now().In(loc),
			PriceDate:     time.Now().In(loc).AddDate(0, 0, -1),
			Open:          2420,
			Close:         2440,
			High:          2440,
			Low:           2415,
			Change:        30,
			ChangePercent: 125,
			Volume:        39270250,
			Amount:        953473542,
		},
	}

	t.Run("valid data", func(t *testing.T) {
		ids := make([]uuid.UUID, len(validPrice))

		// insert test records
		for idx, p := range validPrice {
			id := insertTestRecord(p)
			ids[idx] = id
		}

		result, err := stock.GetStockPriceLatest()

		assert.Nil(t, err)
		assert.Equal(t, 2, len(result))
		assert.Equal(t, "2330", result[0].CompanyID)
		assert.Equal(t, "2454", result[1].CompanyID)

		// delete test records
		for _, id := range ids {
			deleteTestRecord(id)
		}
	})

	t.Run("invalid data", func(t *testing.T) {
		result, err := stock.GetStockPriceLatest()

		assert.Nil(t, err)
		assert.Nil(t, result)
	})
}
