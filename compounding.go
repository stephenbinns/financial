package financial

import "math"

func ConvertAnnualToMonthlyRate(annualRate float64) float64 {
	return math.Pow((annualRate+1.), 1./12.) - 1.
}

func ConvertAnnualToDailyRate(annualRate float64) float64 {
	return math.Pow((annualRate+1.), 1./356.25) - 1.
}

func ConvertMonthlyToAnnual(monthlyRate float64) float64 {
	return math.Pow(monthlyRate+1, 12.) - 1.
}
