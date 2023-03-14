package techan

import "github.com/sdcoffey/big"

type percentIndicator struct {
	indicator Indicator
	base      Indicator
}

// NewFractionIndicator returns a fraction Indicator
func NewFractionIndicator(indicator, base Indicator) Indicator {
	return percentIndicator{indicator: indicator, base: base}
}

func (p percentIndicator) Calculate(index int) big.Decimal {
	return p.indicator.Calculate(index).Div(p.base.Calculate(index))
}
