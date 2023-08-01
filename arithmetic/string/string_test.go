package string

import "testing"

func TestLongestPalindrome(t *testing.T) {
	str := "babc"
	t.Log(longestPalindrome(str))
}

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{
		"faslea",
		"fas",
		"fas",
		"fasl",
		"fasle",
	}
	t.Log(LongestCommonPrefixIII(strs))
}

func TestReverseString(t *testing.T) {
	s := []byte("hello")
	ReverseString(s)
	t.Log(string(s))
}

func TestIsPalindrome(t *testing.T) {
	t.Log(IsPalindrome(""))
}

func TestMyAtoi(t *testing.T) {
	ans := MyAtoi("2147483648")
	t.Log(ans)
}

func TestStrStr(t *testing.T) {
	t.Log(StrStr("mississippi", "issip"))
}

func TestCountAndSay(t *testing.T) {
	t.Log(CountAndSay(10))
}
