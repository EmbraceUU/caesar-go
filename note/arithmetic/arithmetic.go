package arithmetic

import "strconv"

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
你可以假设 nums1 和 nums2 不会同时为空。
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return 0
}

func Reverse(x int) int {
	const IntMax = int(^uint(0)>>32) - 1
	const IntMin = ^IntMax
	var rev int
	for x != 0 {
		pop := x % 10
		x = x / 10
		if rev > IntMax/10 || rev == IntMax/10 && pop > 7 {
			return 0
		}
		if rev < IntMin/10 || rev == IntMin/10 && pop < -8 {
			return 0
		}
		rev = rev*10 + pop
	}
	strconv.Atoi()
	return rev
}
