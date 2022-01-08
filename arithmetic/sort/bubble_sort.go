package sort

func BubbleSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// FIXME: 外层和内层的遍历方向相反了，因为外层是比对的范围，所以从大到小
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				swapI(arr, j, j+1)
			}
		}
	}
}
