package array

// moveZeroes 移动零
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]
// 【数组】【双指针】
// https://leetcode-cn.com/problems/move-zeroes/
func moveZeroes(nums []int) {
	// [0,1,0,3,12]
	// [1,0,3,12,0]
	//	尝试一下归并吧

	// FIXME: 做复杂了，其实直接双指针，一次遍历即可

	var moveMerge = func(nums []int, start, end int) {}
	moveMerge = func(nums []int, start, end int) {
		if start == end {
			return
		}

		// 分治
		if start < end {
			mid := start + (end-start)>>1
			moveMerge(nums, start, mid)
			moveMerge(nums, mid+1, end)

			// [0,1,0,3,12]
			// [0,1,0,    3,12]
			// [0,1,   0,    3,  12]

			// [0,  1,   0,  3,  12]
			// [1,0,  0,   3,12]

			var i, j = start, mid + 1
			if nums[j] == 0 || nums[mid] != 0 {
				return
			}

			for i < len(nums) && nums[i] != 0 {
				i++
			}

			for j < len(nums) && nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j++
			}
		}
	}

	moveMerge(nums, 0, len(nums)-1)
}

func moveZeroesII(nums []int) {
	fast, slow := 0, 0
	// 1,1,0,3,12
	for fast < len(nums) {
		for slow < len(nums) && nums[slow] != 0 {
			slow++
			fast++
		}
		if nums[slow] == 0 && nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}

func moveZeroesIII(nums []int) {
	fast, slow := 0, 0
	// 1,1,0,3,12
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}
