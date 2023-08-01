package array

// https://leetcode.cn/problems/remove-element/
func removeElement(nums []int, val int) int {
	// 0,1,2,2,3,0,4,2
	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
