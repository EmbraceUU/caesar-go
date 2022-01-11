package binary_op

import "testing"

func TestSingleNumberII(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 5, 6, 5, 6, 5, 6, 9}
	num := SingleNumberII(nums)
	t.Log(num)
}
