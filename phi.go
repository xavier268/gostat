package gostat

import "math"

// PHI is the repartition function of the gaussian normal law.
// Algorith adapted from the wikipedia page on normal law :
// https://fr.wikipedia.org/wiki/Loi_normale#D%C3%A9finition_par_la_fonction_de_r%C3%A9partition
func PHI(x float64) float64 {

	if x < -200. { // avoid NaN ...
		return 0.
	}

	if x > 200. { // avoid NaN ...
		return 1.
	}

	var s, t, b, q, i float64

	s = x
	b = x
	q = x * x
	i = 1.

	for math.Abs(s-t) > 1e-10 {
		t = s
		i += 2
		b *= q / i
		s += b
	}
	return 0.5 + s*math.Exp(-0.5*q-0.91893853320467274178)
}

// PHI2 provide the measure between values x and y
// result is always positive.
func PHI2(x, y float64) float64 {
	if y < x {
		x, y = y, x
	}
	return PHI(y) - PHI(x)
}
