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

func TestCombine(t *testing.T) {
	n, k := 1, 1
	result := combine(n, k)
	t.Log(result)
}
