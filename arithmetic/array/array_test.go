package array

import "testing"

func TestTwoSum(t *testing.T) {
	//nums := []int{2, 7, 11, 15}
	//nums := []int{3, 2, 4}
	nums := []int{3, 2}
	target := 6
	t.Log(TwoSum(nums, target))
}

func TestThreeSum(t *testing.T) {
	nums := []int{1, -1, -1, 0}
	t.Log(ThreeSum(nums))
}
