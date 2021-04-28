package binary_tree

import "testing"

func TestNewTreeNode(t *testing.T) {
	arr := [...]int{1, 1}
	nums := make([]int, 0)
	for _, v := range arr {
		nums = append(nums, v)
	}

	root := NewTreeNode(nums)
	IsValidBST(root)
}
