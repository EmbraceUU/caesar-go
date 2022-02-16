package string

import "math"

func reverseInteger(x int) int {
	var rev int
	for x != 0 {
		pop := x % 10
		x = x / 10
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}
