package graph

import "testing"

func TestCloneGraphII(t *testing.T) {
	data := [][]int{
		{2, 4},
		{1, 3},
		{2, 4},
		{1, 3},
	}

	node := NewNode(data)
	cp := CloneGraphIII(node)
	t.Log(cp)
}
