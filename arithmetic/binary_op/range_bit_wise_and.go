package binary_op

// RangeBitwiseAnd 给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）。
// 输入：left = 5, right = 7
// 输出：4
// 输入：left = 0, right = 0
// 输出：0
// 输入：left = 1, right = 2147483647
// 输出：0
// https://leetcode-cn.com/problems/bitwise-and-of-numbers-range/
func RangeBitwiseAnd(left int, right int) int {
	// 如果相等，直接返回
	if left == right {
		return right
	}

	// 如果left为0，说明公共前缀为0
	if left == 0 {
		return 0
	}

	// 直到left和right仅剩公共前缀时，跳出循环
	for left < right {

		// Brian Kernighan 算法
		// 对 number 和 number−1 之间进行按位与运算后，
		// number 中最右边的 1 会被抹去变成 0。
		right &= right - 1
	}
	return right
}
