package techan

import "testing"

func TestFractionIndicator_Calculate(t *testing.T) {
	fi := NewFractionIndicator(NewFixedIndicator(10, 9, 8), NewFixedIndicator(8, 9, 10))

	decimalEquals(t, 1.25, fi.Calculate(0))
	decimalEquals(t, 1, fi.Calculate(1))
	decimalEquals(t, 0.8, fi.Calculate(2))
}
