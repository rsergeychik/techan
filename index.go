package techan

type index struct {
	indicators []Indicator
}

func NewIndex(indicators ...Indicator) *index {
	return &index{
		indicators: indicators,
	}
}

// func (i *index) Calculate(index int) float64 {
// 	var sum float64
// 	for _, indicator := range i.indicators {
// 		sum += indicator.Calculate(index)
// 	}
// 	return sum
// }
