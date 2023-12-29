package stock_price_test

import (
	apperror "stocker-hf-data/web/app/app-error"
	stock "stocker-hf-data/web/app/model/stock_price"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdateStockPrice(t *testing.T) {
	prices := []stock.StockPrice{
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

	testPrep()
	defer testClose()

	ids := make([]uuid.UUID, len(prices))
	// insert test records
	for idx, p := range prices {
		id := insertTestRecord(p)
		ids[idx] = id
	}
	defer func() {
		for _, id := range ids {
			deleteTestRecord(id)
		}
	}()

	// define test cases
	type testcase struct {
		name             string
		price            stock.StockPrice
		expectFail       bool
		expectRowsAffect int64
		expectError      *apperror.ModelError
	}

	subtests := []testcase{
		{
			name: "valid update high price",
			price: stock.StockPrice{
				Uuid:          ids[0].String(),
				CompanyID:     "2330",
				UpdateDate:    time.Now(),
				PriceDate:     time.Now(),
				Open:          50800,
				Close:         50300,
				High:          60000, // update here
				Low:           50000,
				PriceChange:   -700,
				ChangePercent: -139,
				Volume:        19385820,
				Amount:        9753130414,
			},
			expectFail:       false,
			expectRowsAffect: 1,
			expectError:      nil,
		},
		{
			name: "valid update volume",
			price: stock.StockPrice{
				Uuid:          ids[2].String(),
				CompanyID:     "2603",
				UpdateDate:    time.Now(),
				PriceDate:     time.Now().AddDate(0, 0, -1),
				Open:          15100,
				Close:         15050,
				High:          15200,
				Low:           15000,
				PriceChange:   -50,
				ChangePercent: -33,
				Volume:        15427200, // update here
				Amount:        2327960531,
			},
			expectFail:       false,
			expectRowsAffect: 1,
			expectError:      nil,
		},
		{
			name:             "invalid empty price",
			price:            stock.StockPrice{},
			expectFail:       true,
			expectRowsAffect: 0,
			expectError:      apperror.NewModelError(apperror.ErrInputPriceNotValid),
		},
		{
			name: "invalid invalid id",
			price: stock.StockPrice{
				Uuid:          "0000" + ids[3].String()[4:],
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
			expectFail:       true,
			expectRowsAffect: 0,
			expectError:      apperror.NewModelError(apperror.ErrIdNotExist),
		},
	}

	for _, test := range subtests {
		t.Run(test.name, func(t *testing.T) {
			recordAffected, id, err := stock.UpdateStockPrice(test.price)

			if test.expectFail {
				assert.Equal(t, test.expectRowsAffect, recordAffected)
				assert.Equal(t, uuid.UUID{}, id)
				assert.Contains(t, err.Unwrap().Error(), test.expectError.Unwrap().Error())
			} else {
				assert.Equal(t, test.expectRowsAffect, recordAffected)
				assert.NotEqual(t, uuid.UUID{}, id)
				assert.Nil(t, err)
			}
		})
	}

}
