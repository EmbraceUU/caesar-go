package backtrack

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/permutations-ii/
// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
// 输入：nums = [1,1,2]
//	输出：
//	[[1,1,2],
// 	[1,2,1],
// 	[2,1,1]]
func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	track := make([]int, 0)
	used := make([]bool, len(nums))
	// 排序
	sort.Ints(nums)
	btPermuteUnique(&result, track, used, nums)
	return result
}

func btPermuteUnique(result *[][]int, track []int, used []bool, nums []int) {
	// 结果判断
	if len(track) == len(used) {
		temp := make([]int, len(track))
		copy(temp, track)
		*result = append(*result, temp)
		return
	}

	pro := math.MinInt
	// 遍历
	for i := 0; i < len(nums); i++ {
		// 和自己去重
		if used[i] {
			continue
		}
		// 同层去重
		// 注意点：同层去重不能直接用 i和i-1 来比较，因为和自己去重后，同层去重容易误判
		// 	所以选择借助额外变量来记录 pro 的值，可以和自己去重独立开。
		if pro == nums[i] {
			continue
		}

		pro = nums[i]
		used[i] = true
		track = append(track, nums[i])

		btPermuteUnique(result, track, used, nums)

		used[i] = false
		track = track[:len(track)-1]
	}
}
