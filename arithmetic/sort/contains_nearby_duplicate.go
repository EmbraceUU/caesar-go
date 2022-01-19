package sort

// ContainsNearbyDuplicate 存在重复元素 II
// 给你一个整数数组nums 和一个整数k ，判断数组中是否存在两个 不同的索引i和j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。
// 如果存在，返回 true ；否则，返回 false 。
// 输入：nums = [1,2,3,1], k = 3
// 输出：true
// https://leetcode-cn.com/problems/contains-duplicate-ii
func ContainsNearbyDuplicate(nums []int, k int) bool {
	// 借助于map，map的长度为k，
	// 遍历nums的过程中，如果出现重复数字，则返回true
	// 如果map超过了k，则移除最前面的那个num
	nmap := make(map[int]struct{})
	for i, n := range nums {
		if len(nmap) > k {
			delete(nmap, nums[i-k-1])
		}
		_, ok := nmap[n]
		if ok {
			return true
		}
		nmap[n] = struct{}{}
	}
	return false
}
