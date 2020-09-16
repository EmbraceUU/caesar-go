package utils

import (
	"math"
	"testing"
)

func TestAnyToString(t *testing.T) {
	println(AnyToString(math.NaN()), "aaaa")
}
