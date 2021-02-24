package math

import (
	"strconv"
	"testing"
)

func TestPow(t *testing.T) {
	println(strconv.FormatFloat(1200-Pow(2), 'f', -1, 64))
}
