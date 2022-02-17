package string

import "strings"

// isPalindrome 125. 验证回文串
// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
// 说明：本题中，我们将空字符串定义为有效的回文串。
// 【字符串】【双指针】
// https://leetcode-cn.com/problems/valid-palindrome/
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		if !isalnum(s[i]) {
			i++
			continue
		}
		if !isalnum(s[j]) {
			j--
			continue
		}
		if s[i] == s[j] {
			i++
			j--
			continue
		}
		return false
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
