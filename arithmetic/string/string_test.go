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
