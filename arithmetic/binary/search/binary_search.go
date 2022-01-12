package search

// BinarySearch 模板一
// 给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，
// 写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。
// https://leetcode-cn.com/problems/binary-search
func BinarySearch(nums []int, target int) int {
	index := -1

	// 先把空数组排除掉
	if len(nums) == 0 {
		return index
	}

	// 定义min和max
	min, max := 0, len(nums)-1
	// 初始化mid，这样的做法是为了防止溢出
	mid := min + (max-min+1)>>1

	for mid >= min && mid <= max {
		// 如果mid上的数大于target，说明target在左半部分
		// 如果小于target，说明在右半部分
		// 如果相等，说明找到target，直接返回mid
		if nums[mid] > target {
			max = mid - 1
		} else if nums[mid] < target {
			min = mid + 1
		} else {
			return mid
		}

		// 重新计算mid的位置，并进行下一轮查找
		mid = min + (max-min+1)>>1
	}

	return index
}

// BinarySearchII 模板二
// 这个模板更常用一些，可以查找target出现的第一次，最后一次，或者出现的次数统计等。
// 如果没有以上要求，可以直接用模板一。
func BinarySearchII(nums []int, target int) int {
	min, max, mid := 0, len(nums)-1, 0

	// 查找范围至少有3个时才进入循环
	for min+1 < max {
		mid = min + (max-min)>>1

		num := nums[mid]
		switch {
		case num > target:
			max = mid - 1
		case num < target:
			min = mid + 1
		case num == target:
			// 这里没有直接跳出，而是吧mid赋给max
			max = mid
		}
	}

	// 跳出循环后，查找范围少于3个元素，可以直接检查min和max
	if nums[min] == target {
		return min
	}
	if nums[max] == target {
		return max
	}
	return -1
}
