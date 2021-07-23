package gostat

import (
	"fmt"
	"math"
	"testing"
)

func TestPhi(t *testing.T) {

	// z and phi(z)
	data := []float64{
		0, 0.5,
		0.4, 0.6554,
		0.9, 0.8159,
		1.0, 0.8413,
		2.0, 0.9772,
		3., 0.9987,
		3.6, 0.9998,
	}

	for i := 0; i < len(data)-1; i += 2 {
		z, pz := data[i], data[i+1]
		phi := PHI(z)
		if math.Abs(pz-phi) > 1e-4 {
			fmt.Printf("%f\t => %5.4f != %5.4f\n", z, pz, phi)
			t.Fatal()
		}
	}
}

func TestPHI2(t *testing.T) {
	x, y := 0.2, 0.845
	if PHI2(x, y) != PHI2(y, x) {
		t.Fatal("PHI2 should be symetric")
	}
	if PHI2(x, y) <= 0 || PHI2(y, x) <= 0 {
		t.Fatal("PHI2 should be strictly positive !")
	}
	fmt.Println(PHI2(x, y), PHI2(y, x))
}
