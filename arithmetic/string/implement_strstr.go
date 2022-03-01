package string

// strStr 28. 实现 strStr()
// 给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。
// 说明：
//		当needle是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//		对于本题而言，当needle是空字符串时我们应当返回 0 。这与 C 语言的strstr()以及 Java 的indexOf()定义相符。
// https://leetcode-cn.com/problems/implement-strstr
func strStr(haystack string, needle string) int {

	// 时间复杂度：O(m*n)，因为当匹配不符合时，需要返回到初始位置的下一个位置，重新匹配

	if needle == "" {
		return 0
	}

	ans := -1
	if len(needle) > len(haystack) {
		return ans
	}

	var i, j = 0, 0
	for i < len(needle) && j < len(haystack) {
		if haystack[j] != needle[i] {
			// 要恢复
			if i > 0 {
				j = j - i + 1
				i = 0
			} else {
				j++
			}
			continue
		}

		if i == len(needle)-1 {
			return j - i
		}

		i++
		j++
	}

	return ans
}

// strStrKMP
// 【KMP】
func strStrKMP(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}

	// 生成 pi
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}

	// 遍历原字符串
	for i, j := 0, 0; i < n; i++ {
		// 当 原字符串和 匹配字符串的当前位置不相同时，
		// 如果 j > 0，j 跳过一个最大公共前缀的长度 j = pi[i-1]
		for j > 0 && haystack[i] != needle[j] {
			j = pi[i-1]
		}
		// 如果相同，j 往后挪一个位置
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}

	return -1
}
