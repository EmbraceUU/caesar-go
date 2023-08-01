package tree

import "testing"

func TestIsValidBSTII(t *testing.T) {
	nums := []int{2, 2, 2}
	root := NewTreeNode(nums)
	t.Log(isValidBST(root))
}

func TestMaxDepth(t *testing.T) {
	nums := []int{1, -1, 2}
	root := NewTreeNode(nums)
	t.Log(MaxDepth(root))
}

func TestIsSymmetric(t *testing.T) {
	root := NewTreeNode([]int{1, 2, 2, 2, -1, 2})
	t.Log(IsSymmetric(root))
}

func TestSortedArrayToBST(t *testing.T) {
	root := SortedArrayToBST([]int{-10, -3, 0, 5, 9})
	t.Log(root)
}
