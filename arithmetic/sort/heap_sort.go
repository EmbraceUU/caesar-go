package sort

func HeapSort(nums []int) []int {
	// 把nums缔造成一个大根堆
	for i := 0; i < len(nums); i++ {
		heapInsert(nums, i)
	}

	size := len(nums)
	size--
	swap(nums, 0, size)
	for size > 0 {
		heapify(nums, 0, size)
		size--
		swap(nums, 0, size)
	}
	return nums
}

func heapInsert(nums []int, index int) {
	for nums[index] > nums[(index-1)/2] {
		swap(nums, index, (index-1)/2)

		// -1/2=0，所以不用担心数组越界的问题
		index = (index - 1) / 2
	}
}

func heapify(nums []int, index int, size int) {
	left := 2*index + 1
	for left < size {
		// 两个子节点先选出一个大的
		largest := left
		if left+1 < size && nums[left+1] > nums[largest] {
			largest = left + 1
		}

		// 当前已经是大根堆，所以退出循环
		if nums[largest] < nums[index] {
			break
		}

		// 交换节点，并下移index，计算left
		swap(nums, largest, index)
		index = largest
		left = 2*index + 1
	}
}
