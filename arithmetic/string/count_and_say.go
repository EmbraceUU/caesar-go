package string

import (
	"strconv"
	"strings"
)

// countAndSay 38. 外观数列
// 前五项如下：
//
//	1.     1
//	2.     11
//	3.     21
//	4.     1211
//	5.     111221
//	第一项是数字 1
//	描述前一项，这个数是 1 即 “ 一 个 1 ”，记作 "11"
//	描述前一项，这个数是 11 即 “ 二 个 1 ” ，记作 "21"
//	描述前一项，这个数是 21 即 “ 一 个 2 + 一 个 1 ” ，记作 "1211"
//	描述前一项，这个数是 1211 即 “ 一 个 1 + 一 个 2 + 二 个 1 ” ，记作 "111221"
// https://leetcode-cn.com/problems/count-and-say/
func countAndSay(n int) string {
	if n == 0 {
		return "0"
	}

	var counts string
	for i := 1; i <= n; i++ {
		if i == 1 {
			counts = "1"
			continue
		}
		cur := strings.Builder{}
		for j, k := 0, 1; k <= len(counts); k++ {
			if k == len(counts) || counts[j] != counts[k] {
				cur.WriteString(strconv.Itoa(k - j))
				cur.WriteByte(counts[j])
				j = k
			}
		}
		counts = cur.String()
	}
	return counts
}

func countAndSayII(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		cur := strings.Builder{}
		// 同样用到了双指针
		for j, start := 0, 0; j < len(prev); start = j {
			for j < len(prev) && prev[j] == prev[start] {
				j++
			}
			cur.WriteString(strconv.Itoa(j - start))
			cur.WriteByte(prev[start])
		}
		prev = cur.String()
	}
	return prev
}
