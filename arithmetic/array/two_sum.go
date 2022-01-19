package array

// TwoSum 1. 两数之和
// 给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// https://leetcode-cn.com/problems/two-sum
func TwoSum(nums []int, target int) []int {
	// 两层遍历，内外遍历的数值相加判断是否等于target
	// 内层应该比外层范围小，减少重复计算
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
