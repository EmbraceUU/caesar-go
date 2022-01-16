package search

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	target := 6
	t.Log(BinarySearchII(nums, target))
}

func TestRangeSearch(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 3, 4, 5, 6}
	target := 6
	t.Log(RangeSearch(nums, target))
}

func TestMatrixSearch(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	//matrix := [][]int{{1, 3}}
	target := 11
	t.Log(MatrixSearch(matrix, target))
}

func TestFirstBadVersion(t *testing.T) {
	t.Log(FirstBadVersion(1))
}

func TestFindMin(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	t.Log(FindMin(nums))
}

func TestFindMinII(t *testing.T) {
	nums := []int{3, 1, 1}
	t.Log(FindMinII(nums))
}

func TestRotatedSortedArraySearchII(t *testing.T) {
	nums := []int{2, 2, 2}
	var search RotatedSortedArraySearchII
	t.Log(search.search(nums, 1))
}
