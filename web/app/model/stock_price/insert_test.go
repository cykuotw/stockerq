package stock_price_test

import (
	"testing"
	"time"

	apperror "stocker-hf-data/web/app/app-error"
	stock "stocker-hf-data/web/app/model/stock_price"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsertStockPrice(t *testing.T) {
	testPrep()
	defer testClose()

	type testcase struct {
		name             string
		price            stock.StockPrice
		expectFail       bool
		expectRowsAffect int64
		expectError      *apperror.ModelError
	}

	subtests := []testcase{
		{
			name: "valid",
			price: stock.StockPrice{
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
			expectFail:       false,
			expectRowsAffect: 1,
			expectError:      nil,
		},
		{
			name: "valid",
			price: stock.StockPrice{
				CompanyID:     "", // invalid here
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
			expectFail:       true,
			expectRowsAffect: 0,
			expectError:      apperror.NewModelError(apperror.ErrInputPriceNotValid),
		},
	}

	for _, test := range subtests {
		t.Run(test.name, func(t *testing.T) {
			rowsAffected, id, err := stock.InsertStockPrice(test.price)
			defer deleteTestRecord(id)

			if test.expectFail {
				assert.Zero(t, rowsAffected)
				assert.Equal(t, uuid.UUID{}, id)
				assert.NotNil(t, err)
				assert.Equal(t, test.expectError.Unwrap(), err.Unwrap())
			} else {
				assert.Equal(t, test.expectRowsAffect, rowsAffected)
				assert.NotEqual(t, uuid.UUID{}, id)
				assert.Nil(t, err)
			}
		})
	}
}
