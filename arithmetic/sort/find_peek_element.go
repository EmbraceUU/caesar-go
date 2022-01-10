package sort

// FindPeakElement 峰值元素是指其值严格大于左右相邻值的元素。
// 数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
// 你可以假设 nums[-1] = nums[n] = -∞ 。
// 你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
// 输入：nums = [1,2,1,3,5,6,4]
// 输出：1 或 5
// 需要用二分法来处理，才能保证时间复杂度为log n
// https://leetcode-cn.com/problems/find-peak-element/
func FindPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	var min, max = 0, len(nums) - 1
	if nums[min] > nums[min+1] {
		return min
	}
	if nums[max] > nums[max-1] {
		return max
	}

	// 使用二分法
	mid := (max + min) / 2
	for mid >= min && mid <= max {
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}

		if nums[mid] > nums[mid-1] {
			min = mid
			mid = (min + max) / 2
			continue
		}

		max = mid
		mid = (min + max) / 2
	}

	return mid
}
