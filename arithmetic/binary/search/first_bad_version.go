package search

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// FirstBadVersion 你是产品经理，目前正在带领一个团队开发新的产品。
// 不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。
// 假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
// 你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。
// 实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。
//	输入：n = 5, bad = 4
//	输出：4
//	解释：
//	调用 isBadVersion(3) -> false
//	调用 isBadVersion(5) -> true
//	调用 isBadVersion(4) -> true
//	所以，4 是第一个错误的版本。
// https://leetcode-cn.com/problems/first-bad-version
func FirstBadVersion(n int) int {
	// 查找第一次出现的位置，符合二分查找的通用模板
	min, max, mid := 1, n, 0

	for min+1 < max {
		mid = min + (max-min)>>1
		isBad := isBadVersion(mid)
		if isBad {
			max = mid
		} else {
			min = mid + 1
		}
	}

	if isBadVersion(min) {
		return min
	}
	if isBadVersion(max) {
		return max
	}
	return 0
}

func isBadVersion(version int) bool {
	return version == 4
}
