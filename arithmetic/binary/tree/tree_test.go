package tree

import "testing"

func TestIsValidBSTII(t *testing.T) {
	nums := []int{2, 2, 2}
	root := NewTreeNode(nums)
	t.Log(isValidBST(root))
}
