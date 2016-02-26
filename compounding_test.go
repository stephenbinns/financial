package financial

import (
	"testing"
)

func TestCompoundingToMonthly(t *testing.T) {
	rate := ConvertAnnualToMonthlyRate(0.4)

	if rate != 0.028436155726361267 {
		t.Error("Expected 0.028436155726361267, got ", rate)
	}
}

func TestCompoundingToDaily(t *testing.T) {
	rate := ConvertAnnualToDailyRate(0.4)

	if rate != 0.0009449296361851989 {
		t.Error("Expected 0.0009449296361851989, got ", rate)
	}
}

func TestCompoundMonthlyToAnnual(t *testing.T) {
	rate := ConvertMonthlyToAnnual(0.028436155726361267)

	if rate != 0.4000000000000008 {
		t.Error("Expected 0.4, got", rate)
	}
}
