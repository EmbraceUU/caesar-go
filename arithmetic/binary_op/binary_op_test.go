package binary_op

import "testing"

func TestSingleNumberII(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 5, 6, 5, 6, 5, 6, 9}
	num := SingleNumberII(nums)
	t.Log(num)
}

func TestHammingWeight(t *testing.T) {
	var num = uint32(234)
	t.Log(HammingWeight(num))
}

func TestCountBits(t *testing.T) {
	var num = 1
	t.Log(CountBits(num))
}

func TestReverseBits(t *testing.T) {
	var num = uint32(4294967293)
	t.Log(ReverseBits(num))
}
