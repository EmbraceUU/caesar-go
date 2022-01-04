package math

import (
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	z1 := float64(0)
	for math.Abs(z-z1) > 0.000000001 {
		z1 = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}
