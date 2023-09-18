package backtrack

import "sort"

// https://leetcode.cn/problems/combination-sum-ii/
func combinationSum2(candidates []int, target int) [][]int {
	// 和 subsets 的区别是，对 [2], [2,2] 不会有重复
	// 所以在遍历的时候, 对于重复, 只遍历第一个符合条件的元素, 这里的条件, 就是【子集】的条件
	// 1, 2
	// 2| 2
	// 2|

	// 先排序, 让相同元素靠在一起, 如果不排序, 无法去重, 因为 [1,4,4] 和 [4,1,4]是重复的
	sort.Ints(candidates)
	result := make([][]int, 0)
	track := make([]int, 0, len(candidates))
	btCombinationSum2(&result, track, candidates, 0, target, 0)
	return result
}

func btCombinationSum2(result *[][]int, track []int, nums []int, start, target, sum int) {
	// 判断结束状态
	if sum == target {
		temp := make([]int, len(track))
		copy(temp, track)
		*result = append(*result, temp)
		return
	}

	if sum > target {
		return
	}

	for i := start; i < len(nums); i++ {
		// 去重操作
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		track = append(track, nums[i])
		sum += nums[i]

		// 这里需要改成 i+1，而不是 start + 1
		btCombinationSum2(result, track, nums, i+1, target, sum)

		// 错误点：写成 track[:len(nums)-1]
		track = track[:len(track)-1]
		sum -= nums[i]
	}
}
