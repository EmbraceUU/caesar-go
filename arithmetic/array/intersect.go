package array

import "sort"

// intersect 给你两个整数数组nums1 和 nums2 ，请你以数组形式返回两数组的交集。
// 返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。
// 可以不考虑输出结果的顺序。
// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii
func intersect(nums1 []int, nums2 []int) []int {
	// 输入：nums1 = [4,4,5,9],
	//      nums2 = [4,4,4,8,9,9]
	// 输出：[4,9]
	sort.Ints(nums1)
	sort.Ints(nums2)

	var result []int
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			j++
			i++
		} else {
			if nums1[i] > nums2[j] {
				j++
			} else {
				i++
			}
		}
	}

	return result
}
