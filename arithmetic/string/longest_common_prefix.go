package string

// LongestCommonPrefix
// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
// https://leetcode-cn.com/problems/longest-common-prefix/
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var min func(int1, int2 int) int
	min = func(int1, int2 int) int {
		if int1 < int2 {
			return int1
		}
		return int2
	}

	var lcp func(str1, str2 string) string
	lcp = func(str1, str2 string) string {
		l := min(len(str1), len(str2))
		var index int
		for index < l && str1[index] == str2[index] {
			index++
		}
		return str1[:index]
	}

	// 遍历数组的每一个字符串，最终的到最小前缀
	// 时间复杂度：O(mn)
	// 空间复杂度：O(1)

	// 先把第一个字符串作为 result
	// 然后遍历 strs，调用 lcp 方法更新 result

	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result = lcp(result, strs[i])
		if result == "" {
			break
		}
	}
	return result
}

// LongestCommonPrefixII 竖向扫描
func LongestCommonPrefixII(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	//	strs := []string{
	//		"fa",
	//		"fas",
	//		"fasl",
	//		"fasle",
	//	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			// FIXME: 巧妙处理边界问题，当 i 与 len(strs[j]) 相等时，说明 j 上的字符串到头了
			if i == len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// LongestCommonPrefixIII 使用二分查找
func LongestCommonPrefixIII(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 排序得到长度最小的那个
	minL := len(strs[0])
	for _, v := range strs {
		if len(v) < minL {
			minL = len(v)
		}
	}

	isCommonPrefix := func(length int) bool {
		str0, count := strs[0][:length], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}

	// 二分法切割字符串
	start, end := 0, minL
	for start+1 < end {
		mid := start + (end-start)>>1

		// 判断最大前缀的长度的范围
		if isCommonPrefix(mid) {
			start = mid
		} else {
			end = mid
		}
	}

	// 最后再判断一次 start 和 end
	if isCommonPrefix(end) {
		return strs[0][:end]
	}
	return strs[0][:start]
}
