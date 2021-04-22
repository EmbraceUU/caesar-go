package binary_tree

import "testing"

func TestNewTreeNode(t *testing.T) {
	arr := [...]int{8, 4, 9, 3, 5, -1, 10}
	nums := make([]int, 0)
	for _, v := range arr {
		nums = append(nums, v)
	}

	NewTreeNode(nums)
}
