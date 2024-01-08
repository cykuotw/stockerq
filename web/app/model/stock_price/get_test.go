package stock_price_test

import (
	apperror "stockerq/web/app/app-error"
	stock "stockerq/web/app/model/stock_price"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetStockPriceLatest(t *testing.T) {
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

	type testcase struct {
		name            string
		expectFail      bool
		expectCount     int
		expectCompanyID []string
		expectError     *apperror.ModelError
	}

	subtests := []testcase{
		{
			name:            "valid",
			expectFail:      false,
			expectCount:     2,
			expectCompanyID: []string{"2330", "2454"},
			expectError:     nil,
		},
	}

	for _, test := range subtests {
		t.Run(test.name, func(t *testing.T) {
			results, err := stock.GetStockPriceLatest()

			if test.expectFail {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, test.expectCount, len(results))
				for idx, c := range test.expectCompanyID {
					assert.Equal(t, c, results[idx].CompanyID)
				}
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetStockPrice(t *testing.T) {
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

	type testcase struct {
		name            string
		endDate         time.Time
		startDate       time.Time
		expectFail      bool
		expectCount     int
		expectCompanyID []string
		expectError     *apperror.ModelError
	}

	subtests := []testcase{
		{
			name:            "valid",
			startDate:       time.Now(),
			endDate:         time.Now().AddDate(0, 0, 1),
			expectFail:      false,
			expectCount:     2,
			expectCompanyID: []string{"2330", "2454"},
			expectError:     nil,
		},
		{
			name:            "invalid zero time",
			endDate:         time.Time{},
			startDate:       time.Time{},
			expectFail:      true,
			expectCount:     0,
			expectCompanyID: nil,
			expectError:     apperror.NewModelError(apperror.ErrZeroDate),
		},
		{
			name:            "invalid early endDate",
			endDate:         time.Now(),
			startDate:       time.Now(),
			expectFail:      true,
			expectCount:     0,
			expectCompanyID: nil,
			expectError:     apperror.NewModelError(apperror.ErrReverseDate),
		},
	}

	for _, test := range subtests {
		t.Run(test.name, func(t *testing.T) {
			results, err := stock.GetStockPrice(test.startDate, test.endDate)

			if test.expectFail {
				assert.Nil(t, results)
				assert.NotNil(t, err)
				assert.Equal(t, test.expectError.Unwrap(), err.Unwrap())
			} else {
				assert.Equal(t, test.expectCount, len(results))
				for _, p := range results {
					assert.Contains(t, test.expectCompanyID, p.CompanyID)
				}
				assert.Nil(t, err)
			}
		})
	}
}
