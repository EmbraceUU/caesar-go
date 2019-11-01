package hmath

import "math"

// 排列组合
func Combination(m, n int) int {
	if n > m-n {
		n = m - n
	}

	c := 1
	for i := 0; i < n; i++ {
		c *= m - i
		c /= i + 1
	}

	return c
}

func CompareFloat(originData, curData float64, ratioRange float64) bool {
	// 防止一开始原始数据为空
	if originData == 0 {
		return true
	}
	// 如果新数据为空, 不能更新
	if curData == 0 {
		return false
	}
	return math.Abs((curData-originData)/originData) < ratioRange
}
