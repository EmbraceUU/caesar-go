package _decimal

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func StringFix(num float64) {
	fmt.Println(decimal.NewFromFloat(num).StringFixed(-1))
}
