package financial

import (
	"errors"
	"math"
)

func Rate(nper, pmt, pv, fv float64) (float64, error) {
	return RateWithGuess(nper, pmt, pv, fv, 0, 0.1, 1e-6, 100)
}

func RateWithGuess(nper, pmt, pv, fv, when, guess, tol float64, maxIter int) (float64, error) {
	rn := guess
	iter := 0
	close := false

	for iter < maxIter {
		rnp1 := rn - _g_div_gp(rn, nper, pmt, pv, fv, when) // when = 0 for end 1 for beginning
		diff := math.Abs(rnp1 - rn)

		if diff < tol {
			close = true
			break
		}
		iter += 1
		rn = rnp1
	}

	if close == false {
		return 0, errors.New("Cannot calculate rate")
	} else {
		return rn, nil
	}
}

func _g_div_gp(r, n, p, x, y, w float64) float64 {
	t1 := math.Pow((r + 1), n)
	t2 := math.Pow((r + 1), (n - 1))

	return ((y + t1*x + p*(t1-1)*(r*w+1)/r) /
		(n*t2*x - p*(t1-1)*(r*w+1)/math.Pow(r, 2) + n*p*t2*(r*w+1)/r +
			p*(t1-1)*w/r))
}
