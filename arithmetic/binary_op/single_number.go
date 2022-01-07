package binary_op

// SingleNumber 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// https://leetcode-cn.com/problems/single-number/
func SingleNumber(nums []int) int {
	// 10^10=0
	var result int
	for _, v := range nums {
		result = result ^ v
	}
	return result
}
