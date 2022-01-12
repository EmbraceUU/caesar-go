package search

// InsertPosition 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 输入: nums = [1,3,5,6], target = 5
// 输出: 2
func InsertPosition(nums []int, target int) int {
	min, max := 0, len(nums)-1
	var mid int

	for min+1 < max {
		mid = min + (max-min)>>1
		num := nums[mid]
		switch {
		case num > target:
			max = mid - 1
		case num < target:
			min = mid + 1
		case num == target:
			// target只出现了一次，所以直接返回就可以了。
			return mid
		}
	}

	// [min, max]
	// target可能在五个位置，分别是<min，==min，min<target<max, ==max, >max
	switch {
	case nums[min] >= target:
		return min
	case nums[max] == target:
		return max
	case nums[max] < target:
		return max + 1
	}
	return min + 1
}
