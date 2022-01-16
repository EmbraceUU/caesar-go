package search

// RotatedSortedArraySearch 搜索旋转排序数组
type RotatedSortedArraySearch struct{}

// 整数数组 nums 按升序排列，数组中的值 互不相同 。
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，
// 使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
// 例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为[5,6,7,0,1,2,3,4] 。
// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回-1。
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array
func (r RotatedSortedArraySearch) search(nums []int, target int) int {
	// 先用二分法分析数据状况：
	//	1. mid在最小值的右侧或者等于最小值
	//	2. mid在最小值的左侧
	// 然后用mid和max和target做比较
	//	根据mid和max与target值的比较，确定下一次mid移动的方向

	min, max, mid := 0, len(nums)-1, 0

	for min+1 < max {
		mid = min + (max-min)>>1

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < nums[max] {

			if nums[mid] < target && nums[max] >= target {
				min = mid + 1
			} else {
				max = mid - 1
			}

		} else if nums[mid] > nums[max] {

			if nums[mid] > target && nums[max] < target {
				max = mid - 1
			} else {
				min = mid + 1
			}

		}
	}

	if nums[min] == target {
		return min
	} else if nums[max] == target {
		return max
	}
	return -1
}

// RotatedSortedArraySearchII 搜索旋转排序数组 II
type RotatedSortedArraySearchII struct{}

// 同 RotatedSortedArraySearch 题目类似，区别在于 RotatedSortedArraySearchII 中有重复数字
// 思路应该与 RotatedSortedArraySearch 和 FindMinII 相结合
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/
func (r RotatedSortedArraySearchII) search(nums []int, target int) bool {
	min, max, mid := 0, len(nums)-1, 0

	for min+1 < max {
		mid = min + (max-min)>>1
		if nums[mid] == target {
			return true
		}

		if nums[mid] < nums[max] {
			if nums[mid] < target && nums[max] >= target {
				min = mid + 1
			} else {
				max = mid - 1
			}
		} else if nums[mid] > nums[max] {
			if nums[mid] > target && nums[max] < target {
				max = mid - 1
			} else {
				min = mid + 1
			}
		} else {
			max--
		}
	}

	return nums[min] == target || nums[max] == target
}
