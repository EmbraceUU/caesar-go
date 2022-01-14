package search

// FindMin 寻找旋转排序数组中的最小值
// 已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到：
// 若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
// 若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
// 注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
// 给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。
// 请你找出并返回数组中的 最小元素 。
//
//  输入：nums = [4,5,6,7,0,1,2]
//  输出：0
//  解释：原数组为 [0,1,2,4,5,6,7] ，旋转 4 次得到输入数组。
//
// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array
func FindMin(nums []int) int {
	min, max, mid := 0, len(nums)-1, 0
	for min+1 < max {

		mid = min + (max-min)>>1

		// 和找峰值元素「FindPeakElement」有所区别，这里相当于找出峰值中最小的一个位置。
		// 但是二分法同样可以适用。
		// 这里有一个规律，就是nums[mid]>nums[max]时，那说明最小值在mid的右侧
		// 反之则在左侧
		// 如果nums[mid]<nums[mid-1] && nums[mid]<nums[mid+1]，mid则为最小值位置。
		if nums[mid] < nums[mid-1] && nums[mid] < nums[mid+1] {
			return nums[mid]
		}

		// 上面已经判断了mid是否为最小值位置
		// 不再需要把mid包含在下一范围内
		if nums[mid] > nums[max] {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	// 最后的区间内，肯定有一个数是最小数
	if nums[min] > nums[max] {
		return nums[max]
	}
	return nums[min]
}

// FindMinII 寻找旋转排序数组中的最小值 II
// 已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,4,4,5,6,7] 在变化后可能得到：
// 若旋转 4 次，则可以得到 [4,5,6,7,0,1,4]
// 若旋转 7 次，则可以得到 [0,1,4,4,5,6,7]
// 注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
//
// 给你一个可能存在 重复 元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii
func FindMinII(nums []int) int {
	min, max, mid := 0, len(nums)-1, 0
	for min+1 < max {
		mid = min + (max-min)>>1

		// 重复元素的情况，将二分变为三种情况
		// 当mid<max时，最小值应该在mid的左侧或者等于mid
		// 当mid>max时，最小值一定在mid的右侧
		// 当mid==max时，那么mid是可以代替max，所以max--
		if nums[mid] < nums[max] {
			max = mid
		} else if nums[mid] > nums[max] {
			min = mid + 1
		} else {
			max--
		}
	}

	if nums[min] < nums[max] {
		return nums[min]
	}
	return nums[max]
}
