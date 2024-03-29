package techan

import (
	"fmt"
	"strings"

	"github.com/sdcoffey/big"
)

// Candle represents basic market information for a security over a given time period
type Candle struct {
	Period     TimePeriod
	OpenPrice  big.Decimal
	ClosePrice big.Decimal
	MaxPrice   big.Decimal
	MinPrice   big.Decimal
	Volume     big.Decimal
	VWAP       big.Decimal
	TradeCount uint
	totalPrice big.Decimal
}

// NewCandle returns a new *Candle for a given time period
func NewCandle(period TimePeriod) (c *Candle) {
	return &Candle{
		Period:     period,
		OpenPrice:  big.ZERO,
		ClosePrice: big.ZERO,
		MaxPrice:   big.ZERO,
		MinPrice:   big.ZERO,
		VWAP:       big.ZERO,
		Volume:     big.ZERO,
		totalPrice: big.ZERO,
	}
}

// AddTrade adds a trade to this candle. It will determine if the current price is higher or lower than the min or max
// price and increment the tradecount.
func (c *Candle) AddTrade(tradeAmount, tradePrice big.Decimal) {
	if c.OpenPrice.IsZero() {
		c.OpenPrice = tradePrice
	}
	c.ClosePrice = tradePrice

	if c.MaxPrice.IsZero() {
		c.MaxPrice = tradePrice
	} else if tradePrice.GT(c.MaxPrice) {
		c.MaxPrice = tradePrice
	}

	if c.MinPrice.IsZero() {
		c.MinPrice = tradePrice
	} else if tradePrice.LT(c.MinPrice) {
		c.MinPrice = tradePrice
	}

	if c.Volume.IsZero() {
		c.Volume = tradeAmount
	} else {
		c.Volume = c.Volume.Add(tradeAmount)
	}

	if c.totalPrice.IsZero() {
		c.totalPrice = tradePrice.Mul(tradeAmount)
	} else {
		c.totalPrice = c.totalPrice.Add(tradePrice.Mul(tradeAmount))
	}

	if c.VWAP.IsZero() {
		c.VWAP = tradePrice
	} else {
		c.VWAP = c.totalPrice.Div(c.Volume)
	}

	c.TradeCount++
}

func (c *Candle) String() string {
	return strings.TrimSpace(fmt.Sprintf(
		`
Time:	%s
Open:	%s
Close:	%s
High:	%s
Low:	%s
Volume:	%s
VWAP:	%s
	`,
		c.Period,
		c.OpenPrice.FormattedString(2),
		c.ClosePrice.FormattedString(2),
		c.MaxPrice.FormattedString(2),
		c.MinPrice.FormattedString(2),
		c.Volume.FormattedString(2),
		c.VWAP.FormattedString(2),
	))
}
