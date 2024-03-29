package backtrack

import "sort"

// https://leetcode.cn/problems/subsets-ii/
// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//
//	解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
// 输入：nums = [1,2,2]
//	输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
func subsetsWithDup(nums []int) [][]int {
	// 和 subsets 的区别是，对 [2], [2,2] 不会有重复
	// 所以在遍历的时候, 对于重复, 只遍历第一个符合条件的元素, 这里的条件, 就是【子集】的条件
	// 1, 2
	// 2| 2
	// 2|

	// 先排序, 让相同元素靠在一起, 如果不排序, 无法去重, 因为 [1,4,4] 和 [4,1,4]是重复的
	sort.Ints(nums)
	result := make([][]int, 0)
	track := make([]int, 0, len(nums))
	btSubsetsWithDup(&result, track, nums, 0)
	return result
}

func btSubsetsWithDup(result *[][]int, track []int, nums []int, start int) {
	// 判断结束状态
	temp := make([]int, len(track))
	copy(temp, track)
	*result = append(*result, temp)

	for i := start; i < len(nums); i++ {
		// 去重操作
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		track = append(track, nums[i])

		// 这里需要改成 i+1，而不是 start + 1
		btSubsetsWithDup(result, track, nums, i+1)

		// 错误点：写成 track[:len(nums)-1]
		track = track[:len(track)-1]
	}
}
