package search

// RangeSearch 范围查找
func RangeSearch(nums []int, target int) []int {
	min, max, mid := 0, len(nums)-1, 0
	tMin, tMax := -1, -1

	// 范围查找，需要进行两次二分查找过程才能准确定位到
	// 第一次循环是往左半部分找
	// 第二次循环是往右半部分找

	// 第一次二分查找
	for min+1 < max {
		mid = min + (max-min)>>1
		num := nums[mid]

		switch {
		case num > target:
			max = mid - 1
		case num < target:
			min = mid + 1
		case num == target:
			// mid赋值给max，说明找到target后往左半部分继续找
			max = mid
		}
	}

	// 先用min比对target，因为是范围查找，左半部分能找到tMin，
	// 如果min不等于target，再缩小范围判断max是否相等
	// 如果min和max都不相等，说明刚才的循环中也没有找到匹配的位置，
	//（因为target是连贯的），所以直接返回-1即可
	if nums[min] == target {
		tMin = min
	} else if nums[max] == target {
		tMin = max
	} else {
		return []int{-1, -1}
	}

	// 第二次二分查找
	// 这次要在右半部分查找
	min, max = 0, len(nums)-1
	for min+1 < max {
		mid = min + (max-min)>>1
		num := nums[mid]

		switch {
		case num > target:
			max = mid - 1
		case num < target:
			min = mid + 1
		case num == target:
			// 将mid赋值给min
			min = mid
		}
	}

	// 和第一次相反，这次是定位tMax的位置
	if nums[max] == target {
		tMax = max
	} else if nums[min] == target {
		tMax = min
	} else {
		return []int{-1, -1}
	}

	return []int{tMin, tMax}
}
