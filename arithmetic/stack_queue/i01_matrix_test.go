package stack_queue

import "testing"

func TestUpdateMatrix(t *testing.T) {
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	result := UpdateMatrix(mat)
	t.Log(result)
}
