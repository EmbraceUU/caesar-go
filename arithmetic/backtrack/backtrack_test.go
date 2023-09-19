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

func TestSubsetsWithDup(t *testing.T) {
	nums := []int{4, 4, 4, 1, 4}
	result := subsetsWithDup(nums)
	t.Log(result)
	t.Log("[[],[1],[1,4],[1,4,4],[1,4,4,4],[1,4,4,4,4],[4],[4,4],[4,4,4],[4,4,4,4]]")
}

func TestCombinationSum2(t *testing.T) {
	result := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	t.Log(result)
}

func TestPermuteUnique(t *testing.T) {
	result := permuteUnique([]int{1, 1, 2})
	t.Log(result)
}
