package tree

import "testing"

func TestIsValidBSTII(t *testing.T) {
	nums := []int{2, 1, 3}
	root := NewTreeNode(nums)
	t.Log(IsValidBST(root))
}
