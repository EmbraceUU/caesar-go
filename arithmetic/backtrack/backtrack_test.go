package backtrack

import "testing"

func TestSolveNQueens(t *testing.T) {
	result := SolveNQueens(1)
	t.Log(result)
}

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	result := Subsets(nums)
	t.Log(result)
}
