package dynamic

// MinimumTotal 120. 三角形最小路径和
// 给定一个三角形 triangle ，找出自顶向下的最小路径和。
// 每一步只能移动到下一行中相邻的结点上。
// 相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
// 也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
// 输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
// 输出：11
// 解释：如下面简图所示：
//    2
//   3 4
//  6 5 7
// 4 1 8 3
// 自顶向下的最小路径和为11（即，2+3+5+1= 11）。
// https://leetcode-cn.com/problems/triangle
func MinimumTotal(triangle [][]int) int {

	// 自下而上的遍历triangle，没有用到递归和分治法。
	// 使用了一个辅助变量f，用来缓存每个节点上的最小路径和。
	// 这里必须使用辅助变量，不能直接在triangle上做修改，这样会污染原来的节点值。

	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}

	l := len(triangle)
	f := make([][]int, l)

	for i := 0; i < len(triangle); i++ {
		f[i] = make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			f[i][j] = triangle[i][j]
		}
	}

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			f[i][j] = min(f[i+1][j], f[i+1][j+1]) + triangle[i][j]
		}
	}

	return f[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
