package array

// rotate 旋转数组
// 【遍历】
// 最耗时的方法，在LC上会超时
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/
func rotate(nums []int, k int) {
	if len(nums) <= 1 || k == 0 {
		return
	}
	// k 有可能大于或者数倍于 len(nums)
	if k > len(nums) {
		k = k % len(nums)
	}
	var (
		temp int
	)
	for k > 0 {
		temp = nums[len(nums)-1]
		for i := len(nums) - 2; i >= 0; i-- {
			nums[i+1] = nums[i]
		}
		nums[0] = temp
		k--
	}
}

// 【反转】
func rotateII(nums []int, k int) {
	// 先反转整体
	// 再反转前部分
	// 最后反转后部分
	if len(nums) <= 1 || k == 0 {
		return
	}
	// k 有可能大于或者数倍于 len(nums)
	if k > len(nums) {
		k = k % len(nums)
	}
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j := k, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return
}

func rotateIII(nums []int, k int) {
	// temp 3
	// last 2
	// [-1,-100,3,99]
	//2
	if len(nums) <= 1 || k == 0 {
		return
	}
	// k 有可能大于或者数倍于 len(nums)
	if k > len(nums) {
		k = k % len(nums)
	}
	var (
		temp = nums[len(nums)-1]
		last = len(nums) - 1
		end  bool
	)
	for !end {
		target := (last + k) % len(nums)
		temp, nums[target] = nums[target], temp
		last = target

		if last == len(nums)-1 {
			end = true
		}
	}
}
