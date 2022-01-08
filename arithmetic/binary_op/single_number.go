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

// SingleNumberIII 给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。
// https://leetcode-cn.com/problems/single-number-iii/
func SingleNumberIII(nums []int) []int {
	var p int
	// 得到a^b
	for _, v := range nums {
		p ^= v
	}

	// 得到a^b最右边的非0位
	right := p & -p

	var p2 int
	for _, v := range nums {
		// 用非零位的数和每一个元素做 & 运算
		// 将和right非0位同样为1的过滤出来
		// 将这些数做 ^ 运算
		// 得到 a 或者 b
		if v&right == 0 {
			p2 ^= v
		}
	}

	// 返回结果
	return []int{p2, p2 ^ p}
}
