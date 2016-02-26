package financial

import (
	"errors"
	"math"
)

// Fv Returns a float specifying the future value of an annuity based on
// periodic, fixed payments and a fixed interest rate.
func Fv(rate, periods, pmt, principal float64) (float64, error) {
	if periods < 0 {
		return 0, errors.New("Invalid payment period")
	}

	t := math.Pow((rate + 1), float64(periods))
	return ((-principal) * t) - ((pmt / rate) * (t - 1)), nil
}

// Pv Returns a float specifying the present value of an annuity based on periodic,
// fixed payments to be paid in the future and a fixed interest rate.
func Pv(rate, periods, pmt, fv float64, endOfPeriod bool) (float64, error) {
	if rate == 0 {
		return -fv - pmt*periods, nil
	}

	num1 := 1.0
	if endOfPeriod {
		num1 = num1 + rate
	}

	num2 := math.Pow(1.0+rate, periods)
	return -(fv + pmt*num1*((num2-1.0)/rate)) / num2, nil
}

// Pmt Returns a float specifying the payment for an annuity based on periodic,
// fixed payments and a fixed interest rate.
func Pmt(rate, periods, principal, fv float64, endOfPeriod bool) (float64, error) {
	if periods < 1 {
		return 0, errors.New("Invalid payment period")
	}

	if rate == 0 {
		return -(principal / periods), nil
	}

	if endOfPeriod {
		rate = 0
	}

	t := math.Pow((rate + 1), float64(periods))
	return ((-(fv + principal) * t) / (t - 1)) * rate, nil
}

// IPmt Returns a float specifying the interest payment for an annuity based on
// periodic, fixed payments and a fixed interest rate.
func IPmt(rate, currentPeriod, periods, principal float64) (float64, error) {
	if periods < 1 {
		return 0, errors.New("Invalid payment period")
	}
	if currentPeriod > periods {
		return 0, errors.New("Invalid payment period")
	}
	if rate == 0 {
		return 0, nil
	}
	pmt, e := Pmt(rate, periods, principal, 0, false)
	if e != nil {
		return 0, e
	}
	fv, e := Fv(rate, (currentPeriod - 1), pmt, principal)
	if e != nil {
		return 0, e
	}

	return (fv * rate), nil
}

// PPmt Returns a float specifying the principal payment for an annuity based on
// periodic, fixed payments and a fixed interest rate.
func PPmt(rate, currentPeriod, periods, principal float64) (float64, error) {
	pmt, e := Pmt(rate, periods, principal, 0, false)

	if e != nil {
		return 0, e
	}

	ipmt, e := IPmt(rate, currentPeriod, periods, principal)
	if e != nil {
		return 0, e
	}

	return pmt - ipmt, nil
}
