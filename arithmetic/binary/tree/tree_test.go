package tree

import "testing"

func TestIsValidBSTII(t *testing.T) {
	nums := []int{2, 2, 2}
	root := NewTreeNode(nums)
	t.Log(isValidBST(root))
}

func TestIsSymmetric(t *testing.T) {
	root := NewTreeNode([]int{1, 2, 2, 2, -1, 2})
	t.Log(IsSymmetric(root))
}
