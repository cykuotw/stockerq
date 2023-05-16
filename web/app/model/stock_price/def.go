package stock_price

import (
	"time"

	"github.com/google/uuid"
)

type StockPrice struct {
	Id            uuid.UUID `json:"id"`
	CompanyID     string    `json:"company_id"`
	UpdateDate    time.Time `json:"update_date"`
	PriceDate     time.Time `json:"price_date"`
	Open          int32     `json:"open"`           // in $0.01 NTD
	Close         int32     `json:"close"`          // in $0.01 NTD
	High          int32     `json:"high"`           // in $0.01 NTD
	Low           int32     `json:"low"`            // in $0.01 NTD
	Change        int64     `json:"change"`         // in $0.01 NTD
	ChangePercent int64     `json:"change_percent"` // in 0.01%
	Volume        int64     `json:"volume"`         // in share
	Amount        int64     `json:"amount"`         // in NTD
}

func (price *StockPrice) isValid() bool {
	result := true

	result = result && (len(price.CompanyID) >= 4)
	result = result && (!price.UpdateDate.IsZero())
	result = result && (!price.PriceDate.IsZero())
	result = result && (price.Open > 0)
	result = result && (price.Close > 0)
	result = result && (price.High > 0)
	result = result && (price.Low > 0)
	result = result && (price.ChangePercent-int64((float64(price.Change)/float64(price.Close))*10000) < 2)
	result = result && (price.Volume >= 0)
	result = result && (price.Amount >= 0)

	return result
}
