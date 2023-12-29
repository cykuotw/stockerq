package stock_price_test

import (
	stock "stocker-hf-data/web/app/model/stock_price"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetStockPriceLatest(t *testing.T) {
	testPrep()
	defer testClose()

	validPrice := []stock.StockPrice{
		{
			CompanyID:     "2330",
			UpdateDate:    time.Now(),
			PriceDate:     time.Now(),
			Open:          50800,
			Close:         50300,
			High:          50800,
			Low:           50000,
			PriceChange:   -700,
			ChangePercent: -139,
			Volume:        19385820,
			Amount:        9753130414,
		},
		{
			CompanyID:     "2454",
			UpdateDate:    time.Now(),
			PriceDate:     time.Now(),
			Open:          69100,
			Close:         69800,
			High:          69800,
			Low:           69000,
			PriceChange:   900,
			ChangePercent: 131,
			Volume:        5282896,
			Amount:        3673257462,
		},
		{
			CompanyID:     "2603",
			UpdateDate:    time.Now(),
			PriceDate:     time.Now().AddDate(0, 0, -1),
			Open:          15100,
			Close:         15050,
			High:          15200,
			Low:           15000,
			PriceChange:   -50,
			ChangePercent: -33,
			Volume:        15427176,
			Amount:        2327960531,
		},
		{
			CompanyID:     "2891",
			UpdateDate:    time.Now(),
			PriceDate:     time.Now().AddDate(0, 0, -1),
			Open:          2420,
			Close:         2440,
			High:          2440,
			Low:           2415,
			PriceChange:   30,
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
		if len(result) != 0 {
			assert.Equal(t, "2330", result[0].CompanyID)
			assert.Equal(t, "2454", result[1].CompanyID)
		}

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

func TestGetStockPrice(t *testing.T) {
	testPrep()
	defer testClose()

	validPrice := []stock.StockPrice{
		{
			CompanyID:     "2330",
			UpdateDate:    time.Now(),
			PriceDate:     time.Now(),
			Open:          50800,
			Close:         50300,
			High:          50800,
			Low:           50000,
			PriceChange:   -700,
			ChangePercent: -139,
			Volume:        19385820,
			Amount:        9753130414,
		},
		{
			CompanyID:     "2454",
			UpdateDate:    time.Now(),
			PriceDate:     time.Now(),
			Open:          69100,
			Close:         69800,
			High:          69800,
			Low:           69000,
			PriceChange:   900,
			ChangePercent: 131,
			Volume:        5282896,
			Amount:        3673257462,
		},
		{
			CompanyID:     "2603",
			UpdateDate:    time.Now(),
			PriceDate:     time.Now().AddDate(0, 0, 2),
			Open:          15100,
			Close:         15050,
			High:          15200,
			Low:           15000,
			PriceChange:   -50,
			ChangePercent: -33,
			Volume:        15427176,
			Amount:        2327960531,
		},
		{
			CompanyID:     "2891",
			UpdateDate:    time.Now(),
			PriceDate:     time.Now().AddDate(0, 0, 2),
			Open:          2420,
			Close:         2440,
			High:          2440,
			Low:           2415,
			PriceChange:   30,
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

		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		tomorrow := today.AddDate(0, 0, 1)
		result, err := stock.GetStockPrice(today, tomorrow)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(result))
		if len(result) != 0 {
			assert.Equal(t, "2330", result[0].CompanyID)
			assert.Equal(t, "2454", result[1].CompanyID)
		}

		// delete test records
		for _, id := range ids {
			deleteTestRecord(id)
		}
	})

	t.Run("invalid data", func(t *testing.T) {
		// zero time
		result, err := stock.GetStockPrice(time.Time{}, time.Time{})
		assert.Nil(t, result)
		assert.NotNil(t, err)

		result, err = stock.GetStockPrice(time.Now(), time.Time{})
		assert.Nil(t, result)
		assert.NotNil(t, err)

		result, err = stock.GetStockPrice(time.Time{}, time.Now())
		assert.Nil(t, result)
		assert.NotNil(t, err)

		// early endDate
		now := time.Now()
		result, err = stock.GetStockPrice(now, now)
		assert.Nil(t, result)
		assert.NotNil(t, err)

		result, err = stock.GetStockPrice(now, now)
		assert.Nil(t, result)
		assert.NotNil(t, err)
	})
}
