package sort

/**
选择排序
*/

func SelectionSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// 从0~N-1遍历
	for i := 0; i < len(arr)-1; i++ {
		// 先设定i位置为当前最小位置
		p := i
		for j := i + 1; j < len(arr); j++ {
			// 如果p位置比j位置大, 将j位置设为最小值位置
			if arr[p] > arr[j] {
				p = j
			}
		}
		// 将i位置和p位置交换
		swapI(arr, i, p)
	}
}

func swapI(arr []int, i, j int) {
	if i == j {
		return
	}
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}
