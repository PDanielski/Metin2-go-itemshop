package item

import (
	"mt2is/pkg/wallet"
	"time"
)

type Discount struct {
	ID                 int
	ItemID             int
	Currency           *wallet.Currency
	StartTime          time.Time
	EndTime            time.Time
	AbsoluteModifier   int
	RelativeModifier   int
	PercentageModifier float32
}
