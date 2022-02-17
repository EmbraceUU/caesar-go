package string

// isAnagram 242. 有效的字母异位词
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。
// 输入: s = "anagram", t = "nagaram"
// 输出: true
// https://leetcode-cn.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	// 先判断是否长度一致
	if len(s) != len(t) {
		return false
	}
	// 定义map，存放字符和出现的次数
	m := make(map[int32]int)
	// 遍历s将字符和次数存入map
	for _, v := range s {
		d := v - 'a'
		m[d]++
	}
	// 遍历t将map逐步清空
	for _, v := range t {
		d := v - 'a'
		m[d]--
		if m[d] < 0 {
			return false
		}
	}
	// 因为长度一致，所以如果不符合条件，必有一个位置是出现多次的，遍历t时出现小于0的情况，直接返回了。
	return true
}
