package sort

// 【归并排序 分治法】
func MergeSort(nums []int) []int {
	// 如果只有1个或者没有元素，直接返回，不需要再排序
	if len(nums) <= 1 {
		return nums
	}

	// 分解
	mid := len(nums) / 2
	// TODO 这里截取出错了，冒号前后是左闭右开的
	nums1 := nums[:mid]
	nums2 := nums[mid:]
	sortednums1 := MergeSort(nums1)
	sortednums2 := MergeSort(nums2)

	// 合并
	numsNew := make([]int, 0)

	for len(sortednums1) > 0 && len(sortednums2) > 0 {
		if sortednums1[0] <= sortednums2[0] {
			numsNew = append(numsNew, sortednums1[0])
			sortednums1 = sortednums1[1:]
			// TODO 这里出错了, 报了out of range
		} else if sortednums1[0] > sortednums2[0] {
			numsNew = append(numsNew, sortednums2[0])
			sortednums2 = sortednums2[1:]
		}
	}

	numsNew = append(numsNew, sortednums1...)
	numsNew = append(numsNew, sortednums2...)

	return numsNew
}
