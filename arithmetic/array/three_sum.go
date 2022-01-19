package array

import "sort"

// ThreeSum 三数之和
//
// 给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
//
// 重点是不重复，如果允许重复，可以直接使用类似 TwoSum 的思路，三层循环遍历即可。
// 去重需要用到 排序 + 双指针 的技巧
// https://leetcode-cn.com/problems/3sum
func ThreeSum(nums []int) [][]int {

	// 先对数组进行排序
	// 遍历过程中，对相邻两个num重复的，直接跳过，这样即可去重
	// 然后将第二层循环和第三层循环合并，使用双指针进行遍历，第二层从左往右遍历，第三层在第二层从右往左遍历。
	// 因为 a + b + c = 0，并且数组是递增且通过逻辑避免重复的，所以 a + b1 + c1 = 0，b1 > b && c1 < c。

	var res [][]int

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {

		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		k := len(nums) - 1

		for j := i + 1; j < len(nums) && k > j; j++ {

			// nums[j] == nums[j-1]，用来防止重复
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}

			// 一直减出现数组越界了
			for k > 0 && nums[i]+nums[j]+nums[k] > 0 {
				k--
			}

			// k > j，表示第二层循环必须在第三层的左边，这样防止重复
			if k <= j {
				break
			}

			if nums[i]+nums[j]+nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
			} else {
				// 跳过此次循环，但是不能break，继续让第二层增加。
				continue
			}
		}
	}
	return res
}
