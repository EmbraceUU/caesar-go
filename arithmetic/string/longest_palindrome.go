package string

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		// 以 s[i] 为中心的最长回文子串
		s1 := palindrome(s, i, i)
		// 以 s[i] 和 s[i+1] 为中心的最长回文子串
		s2 := palindrome(s, i, i+1)
		// res = longest(res, s1, s2)
		if len(res) > len(s1) {
			res = res
		} else {
			res = s1
		}
		if len(res) > len(s2) {
			res = res
		} else {
			res = s2
		}
	}
	return res
}

func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
