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

// SingleNumberII 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
// https://leetcode-cn.com/problems/single-number-ii/
func SingleNumberII(nums []int) int {
	var result int
	// 假设int是64位
	for i := 0; i < 64; i++ {
		// 定义一个sum，表示每个数当前位的数值的总和
		var sum int

		// 遍历每个num
		for j := 0; j < len(nums); j++ {
			// [num & 1]可以得到最后一位的值
			// 通过[num >> i]可以遍历到每一位
			// 然后将每个num的相同位相加
			sum += (nums[j] >> i) & 1
		}

		// 题目是其他所有数都是3个
		// 那么每一位的[sum%3]得到出现一次的num的i位置的数值
		// 通过[num << i]再与result求与，即可恢复result在i位置的数值
		result |= (sum % 3) << i
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
