package ext

import (
	"testing"
)

func TestCalcCost(t *testing.T) {
	msg := CalcCost(2, 4, 2)
	if msg != 4 {
		t.Fatalf(`Expected 4 instead %v`, msg)
	}
}
