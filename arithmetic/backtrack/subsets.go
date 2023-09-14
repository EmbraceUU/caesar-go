package backtrack

func Subsets(nums []int) [][]int {
	return subsets(nums)
}

// https://leetcode.cn/problems/subsets/description/
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	track := make([]int, 0)
	btSubsets(&result, track, nums, 0)
	return result
}

// 核心点：如何判断已经选择的路径 ？
// 比如已经有过 1，2 就不能再有 2，1
// 加上一个限定条件，内部循环只能从当前索引开始向后找
// 去重逻辑，可以使用一个 col 变量，只要从 col 后面开始遍历，就不会出现重复的情况

// 我们通过保证元素之间的相对顺序不变来防止出现重复的子集。

func btSubsets(result *[][]int, track []int, nums []int, start int) {
	// 判断结束状态
	temp := make([]int, len(track))
	copy(temp, track)
	*result = append(*result, temp)

	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])

		// 这里需要改成 i+1，而不是 start + 1
		btSubsets(result, track, nums, i+1)

		track = track[:len(track)-1]
	}
}

func subsets2(nums []int) [][]int {

	var res [][]int
	var track []int

	// 定义回溯函数
	var backtrack func(start int)
	backtrack = func(start int) {
		// 每次加入一个子集
		temp := make([]int, len(track))
		copy(temp, track)
		res = append(res, temp)

		// 回溯算法框架
		for i := start; i < len(nums); i++ {
			// 做出选择
			track = append(track, nums[i])

			// 递归进入下一个状态
			backtrack(i + 1)

			// 撤销上一个状态
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return res
}
