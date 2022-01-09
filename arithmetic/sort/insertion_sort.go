package sort

func InsertionSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// 插入排序, 是从0位置开始遍历，比较0~i位置的排序情况，i位置往0位置遍历，碰到比自己大的元素进行交换，碰到比自己小的，退出循环
	for i := range arr {
		if i == 0 {
			continue
		}

		for j := i; j > 0; j-- {
			switch {
			case arr[j] > arr[j-1]:
				break
			case arr[j] == arr[j-1]:
				continue
			case arr[j] < arr[j-1]:
				swapI(arr, j, j-1)
			}
		}
	}
}
