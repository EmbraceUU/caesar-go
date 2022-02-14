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
