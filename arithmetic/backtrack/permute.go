package backtrack

// https://leetcode.cn/problems/permutations/
// 利用回溯的解题结构来做
func permute(nums []int) [][]int {
	var result [][]int
	var track []int
	used := make([]bool, len(nums))

	bt(nums, track, used, &result)
	return result
}

func bt(nums []int, track []int, used []bool, result *[][]int) {
	// if 满足结束条件:
	//    result.add(路径)
	//    return
	//
	// for 选择 in 选择列表:
	//    做选择
	//    backtrack(路径, 选择列表)
	//    撤销选择

	// 判断是否满足结束条件
	if len(track) == len(nums) {
		temp := make([]int, len(track))
		copy(temp, track)
		*result = append(*result, temp)
		return
	}

	for i := range nums {
		// 避免重复选择/往回走
		if used[i] {
			continue
		}

		// 做选择
		track = append(track, nums[i])
		used[i] = true
		bt(nums, track, used, result)
		// 撤销选择
		track = track[:len(track)-1]
		used[i] = false
	}
}
