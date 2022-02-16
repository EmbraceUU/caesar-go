package string

import "testing"

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
