package op

// HammingWeight 编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。
// https://leetcode-cn.com/problems/number-of-1-bits/
func HammingWeight(num uint32) int {
	var sum int

	// 方法一
	// 遍历num的每一位
	for i := 0; i < 32; i++ {
		// 通过[num&1]求出最右位置的值
		// 通过右移遍历每一位
		sum += int((num >> i) & 1)
	}

	sum = 0

	// 方法二
	// 当num>0时，说明num中还存在1
	// n & (n−1)，其运算结果恰为把 n 的二进制位中的最低位的 1 变为 0 之后的结果
	// 6 & 6-1 -> 6(110) & 5(101) -> 4(100)
	for ; num > 0; num &= num - 1 {
		sum++
	}
	return sum
}

// CountBits 给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。
// https://leetcode-cn.com/problems/counting-bits/
func CountBits(n int) []int {
	var ans []int
	ans = append(ans, 0)

	for i := 1; i <= n; i++ {
		var sum int
		j := i
		for ; j > 0; j &= j - 1 {
			sum++
		}
		ans = append(ans, sum)
	}

	return ans
}
