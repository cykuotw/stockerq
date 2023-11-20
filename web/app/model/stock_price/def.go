package stock_price

import (
	"time"
)

type StockPrice struct {
	Uuid          string    `json:"uuid"`
	CompanyID     string    `json:"company_id"`
	UpdateDate    time.Time `json:"update_date"`
	PriceDate     time.Time `json:"price_date"`
	Open          uint32    `json:"open"`           // in $0.01 NTD
	Close         uint32    `json:"close"`          // in $0.01 NTD
	High          uint32    `json:"high"`           // in $0.01 NTD
	Low           uint32    `json:"low"`            // in $0.01 NTD
	PriceChange   int64     `json:"price_change"`   // in $0.01 NTD
	ChangePercent int64     `json:"change_percent"` // in 0.01%
	Volume        uint64    `json:"volume"`         // in share
	Amount        uint64    `json:"amount"`         // in NTD
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
	result = result && (price.ChangePercent-int64((float64(price.PriceChange)/float64(price.Close))*10000) < 2)
	result = result && (price.Volume >= 0)
	result = result && (price.Amount >= 0)

	return result
}
