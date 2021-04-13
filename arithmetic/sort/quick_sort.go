package sort

// QuickSort 快速排序 分治法
func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	// 当分区内只有一个元素时, 不再进行比较
	if start < end {
		// 分治法
		i := partition(nums, start, end)
		quickSort(nums, start, i-1)
		quickSort(nums, i+1, end)
	}
}

func partition(nums []int, start, end int) int {
	// 只保证分区, 不保证两侧的顺序
	p := nums[end]
	i := nums[start]
	// 遍历一遍, 比p小的往前换, 比p大的往后换
	for j := start; j < end; j++ {
		if nums[j] < p {
			swap(nums, i, j)
			// 每更换一次, i向后移动一下
			i++
		}
	}
	// 将p移动到中间, 左边都是小的, 右边都是大的
	swap(nums, i, end)
	return i
}

func swap(nums []int, i, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}
