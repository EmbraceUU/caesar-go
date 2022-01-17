package sort

import "math/rand"

// QuickSortII 快速排序II
// 使用随机选取pivot的方法
// 使用随机选取的方法，时间复杂度的期望是O(N*logN)的
func QuickSortII(nums []int) {
	quickSortII(nums, 0, len(nums)-1)
}

func quickSortII(nums []int, l, r int) {
	if l < r {
		// 对l，r范围内进行分区
		i := randomPartition(nums, l, r)
		// 递归调用，对两侧的区间再进行分区
		quickSortII(nums, l, i-1)
		quickSortII(nums, i+1, r)
	}
}

func randomPartition(nums []int, l, r int) int {
	// 随机选取一个数，与r位置进行交换
	// 然后再进行分区
	random := rand.Intn(r - l + 1)
	swap(nums, r, l+random)
	return partitionII(nums, l, r)
}

func partitionII(nums []int, l, r int) int {
	// 声明三个索引，i表示<pivot的区域，j表示当前遍历到的位置，p是pivot的位置
	i, j, p := l-1, l, r

	// j在p的前面遍历，当j和p重合时，跳出遍历
	for j < p {

		// 当j位置的数值<p位置的数值时
		// 先把i往后移动一位，然后和j交换
		if nums[j] < nums[p] {
			i++
			swap(nums, i, j)
		}
		j++
	}

	// 最后，把p和i后面一位做交换
	// 这样，区间l和r就成了 (l~i-1) i (i+1, r)，并且i左边都比i小，i都比i大
	i++
	swap(nums, i, j)
	return i
}
