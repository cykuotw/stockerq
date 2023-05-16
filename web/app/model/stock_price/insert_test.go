package stock_price_test

import (
	"testing"
	"time"

	stock "stocker-quant/web/app/model/stock_price"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsertStockPrice(t *testing.T) {
	testPrep()
	defer testClose()

	priceValid := stock.StockPrice{
		CompanyID:     "2330",
		UpdateDate:    time.Now(),
		PriceDate:     time.Now(),
		Open:          50800,
		Close:         50300,
		High:          50800,
		Low:           50000,
		Change:        -700,
		ChangePercent: -139,
		Volume:        19385820,
		Amount:        9753130414,
	}
	t.Run("valid parameter", func(t *testing.T) {
		rowsAffected, id, err := stock.InsertStockPrice(priceValid)

		assert.Nil(t, err)
		assert.NotEqual(t, uuid.UUID{}, id)
		assert.Equal(t, int64(1), rowsAffected)
	})
	t.Run("invalid parameter", func(t *testing.T) {
		priceInvalid := priceValid
		priceInvalid.CompanyID = ""
		rowsAffected, id, err := stock.InsertStockPrice(priceInvalid)

		assert.NotNil(t, err)
		assert.Equal(t, uuid.UUID{}, id)
		assert.Zero(t, rowsAffected)
	})
}
