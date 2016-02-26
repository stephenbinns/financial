package financial

import (
	"math"
	"testing"
)

func TestCalculatesRate(t *testing.T) {
	rate := 0.07
	pv := -100.
	term := 60.
	pmt, _ := Pmt(rate, term, pv, 0, false)
	actual, err := Rate(term, pmt, pv, 0)

	if err != nil {
		t.Error("Expected no error")
	}
	if rate != Round(actual, 3) {
		t.Error("Expected: 0.07 but was", Round(actual, 3))
	}
}

func TestCalculatesRateWithGuess(t *testing.T) {
	rate := 0.07
	pv := -100.
	term := 60.
	pmt, _ := Pmt(rate, term, pv, 0, false)
	actual, err := RateWithGuess(term, pmt, pv, 0, 0, 0.08, 1e-6, 50)

	if err != nil {
		t.Error("Expected no error")
	}
	if rate != Round(actual, 3) {
		t.Error("Expected: 0.07 but was", Round(actual, 3))
	}
}

func Round(x float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(x*shift) / shift
}
