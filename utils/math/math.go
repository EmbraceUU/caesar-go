package math

import "math"

func Pow(prec int) float64 {
	pow := math.Pow10(-prec) * 5
	return pow
}
