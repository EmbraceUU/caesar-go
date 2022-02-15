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

func TestRotate(t *testing.T) {
	nums := []int{-1, -100, 3, 99}
	k := 2
	Rotate(nums, k)
	t.Log(nums)
}

func TestPlusOne(t *testing.T) {
	nums := []int{9, 9, 9, 9, 9}
	result := PlusOne(nums)
	t.Log(result)
}

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	MoveZeroes(nums)
	t.Log(nums)
}
