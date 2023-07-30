package array

// removeDuplicates 删除数组中的重复项
// 【双指针】，【数组】
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	var (
		k int
	)

	// [0,1,2,3,4k,2,2,3,3,4]i

	if len(nums) == 0 {
		return k
	}
	for i := 0; i < len(nums); i++ {
		if nums[k] != nums[i] {
			nums[k+1] = nums[i]
			k++
		}
	}

	return k + 1
}

func RemoveDuplicatesII(nums []int) int {
	// 快慢指针
	// p1 是唯一元素的个数
	// p2 是当前遍历的进度
	// 1，2，3，4，5，5，5，6
	p := 0
	for i := 1; i < len(nums); i++ {
		// 如果一样，i++
		// 如果不一样，p++，i位置覆盖到p位置
		if nums[p] == nums[i] {
			continue
		}

		p++
		if p != i {
			nums[p] = nums[i]
		}
	}
	p++
	return p
}
