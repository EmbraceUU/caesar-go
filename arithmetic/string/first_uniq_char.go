package string

// firstUniqChar 字符串中的第一个唯一字符
// 给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。
// https://leetcode-cn.com/problems/first-unique-character-in-a-string/
func firstUniqChar(s string) int {
	// 如果借用了队列，也是省去了索引排序的过程
	// 定义一个 map，存放字符和索引
	n := len(s)
	m := [26]int{}
	for i := range m {
		m[i] = n
	}

	for i, v := range s {
		d := v - 'a'
		if m[d] == n {
			// 第一次出现设置为索引
			m[d] = i
		} else {
			// 重复出现的字符 value 为-1
			m[d] = -1
		}
	}
	// 将索引最小的返回
	ans := n
	for _, v := range m {
		if v == n || v == -1 {
			continue
		}
		if v < ans {
			ans = v
		}
	}

	if ans < n {
		return ans
	}
	return -1
}
