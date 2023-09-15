package backtrack

// https://leetcode.cn/problems/combinations/
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
// 组合，不可重复
func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	// 失误点：sub的容量没有用k变量来控制
	sub := make([]int, 0, k)
	btCombine(&result, sub, 1, n)
	return result
}

func btCombine(result *[][]int, sub []int, start, n int) {
	// 检查是否满足结果，控制递归是否继续
	if len(sub) == cap(sub) {
		temp := make([]int, len(sub))
		copy(temp, sub)
		*result = append(*result, temp)
		return
	}

	// 通过 start 来控制向下&向右遍历，不会有重复
	for i := start; i <= n; i++ {
		sub = append(sub, i)

		btCombine(result, sub, i+1, n)

		sub = sub[:len(sub)-1]
	}
}
